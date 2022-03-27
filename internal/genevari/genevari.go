//go:build generics

package genevari

import (
	"github.com/MistaTwista/generigo/internal/util"
)

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
