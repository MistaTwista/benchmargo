package genevari

func Pipe(fns ...func(int) int) int {
	var r = int(0)
	for _, fn := range fns {
		r = fn(r)
	}

	return r
}

func Identity(n int) func(int) int {
	return func(_ int) int {
		return n
	}
}

func Multiply(i int) int {
	return i * i
}

func Plus(i int) int {
	return i + i
}

func MultiplyWith(n int) func(int) int {
	return func(i int) int {
		return i * n
	}
}

func Generator(num int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for i := 0; i < num; i++ {
			res <- i
		}
	}()

	return res
}

func Multiplicator(mulBy int) func(<-chan int) <-chan int {
	return func(numCh <-chan int) <-chan int {
		res := make(chan int)
		go func() {
			defer close(res)
			for n := range numCh {
				res <- n * mulBy
			}
		}()
		return res
	}
}

func Duplicator(numCh <-chan int) <-chan []int {
	res := make(chan []int)
	go func() {
		defer close(res)
		for n := range numCh {
			res <- []int{n, n}
		}
	}()
	return res
}

func Summer(numCh <-chan []int) chan int {
	res := make(chan int)
	go func() {
		var sum int
		defer close(res)
		defer func() {
			res <- sum
		}()

		for list := range numCh {
			for _, n := range list {
				sum += n
			}
		}
	}()

	return res
}