package genepipe

func Generator(num int, seeds ...int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for i := 0; i < num; i++ {
			for _, s := range seeds {
				res <- s
			}
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

func Flat(numCh <-chan []int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for list := range numCh {
			for _, n := range list {
				res <- n
			}
		}
	}()
	return res
}

func Add(numCh <-chan int) chan int {
	res := make(chan int)
	go func() {
		var sum int
		defer close(res)
		defer func() {
			res <- sum
		}()

		for n := range numCh {
			sum += n
		}
	}()

	return res
}
