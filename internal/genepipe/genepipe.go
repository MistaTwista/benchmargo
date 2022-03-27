//go:build generics

package genepipe

import "github.com/MistaTwista/benchmargo/internal/util"

func GeneGenerator[T any](num int, seeds ...T) <-chan T {
	res := make(chan T)
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

func GeneDuplicator[T any](numCh <-chan T) <-chan []T {
	res := make(chan []T)
	go func() {
		defer close(res)
		for n := range numCh {
			res <- []T{n, n}
		}
	}()
	return res
}

func GeneFlat[T any](numCh <-chan []T) <-chan T {
	res := make(chan T)
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

func GeneSummer[T util.Addable](numCh <-chan T) chan T {
	res := make(chan T)
	go func() {
		var s T
		defer close(res)
		defer func() {
			res <- s
		}()

		for n := range numCh {
			s = s + n
		}
	}()

	return res
}
