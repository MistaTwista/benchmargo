package genevari

import (
	"github.com/MistaTwista/generigo/internal/util"
)

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

var Multiply = func(i int) int {
	return i * i
}

var Plus = func(i int) int {
	return i + i
}

func MultiplyWith(n int) func(int) int {
	return func(i int) int {
		return i * n
	}
}
func GPipe[T util.Numbers](fns ...func(T) T) T {
	var r = T(0)
	for _, fn := range fns {
		r = fn(r)
	}

	return r
}

func GIdentity[T util.Numbers](n T) func(T) T {
	return func(_ T) T {
		return n
	}
}

func GMultiply[T util.Numbers](i T) T {
	return i * i
}

func GPlus[T util.Numbers](i T) T {
	return i + i
}

func GMultiplyWith[T util.Numbers](n T) func(T) T {
	return func(i T) T {
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

func GeneGenerator[T util.Numbers](num int) <-chan T {
	res := make(chan T)
	go func() {
		defer close(res)
		for i := 0; i < num; i++ {
			res <- T(i)
		}
	}()

	return res
}

func GeneMultiplicator[T util.Numbers](mulBy T) func(<-chan T) <-chan T {
	return func(numCh <-chan T) <-chan T {
		res := make(chan T)
		go func() {
			defer close(res)
			for n := range numCh {
				res <- n * mulBy
			}
		}()
		return res
	}
}

func GeneDuplicator[T util.Numbers](numCh <-chan T) <-chan []T {
	res := make(chan []T)
	go func() {
		defer close(res)
		for n := range numCh {
			res <- []T{n, n}
		}
	}()
	return res
}

func GeneSummer[T util.Numbers](numCh <-chan []T) chan T {
	res := make(chan T)
	go func() {
		var s T
		defer close(res)
		defer func() {
			res <- s
		}()

		for list := range numCh {
			for _, n := range list {
				s += n
			}
		}
	}()

	return res
}
