//go:build generics

package genetochan

import "sync"

func GToChanAny[T any, R func() T](fn R) <-chan T {
	res := make(chan T)
	go func() {
		defer func() {
			close(res)
		}()

		x := fn()
		res <- x
	}()

	return res
}

func GToChan[T any](fn func() T) <-chan T {
	res := make(chan T)
	go func() {
		defer func() {
			close(res)
		}()

		x := fn()
		res <- x
	}()

	return res
}

func GMix[T any](chans ...<-chan T) chan T {
	res := make(chan T)
	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)
		go func(dataCh <-chan T) {
			defer wg.Done()
			res <- <-dataCh
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func With[T any](some T, fn func(T) T) func() T {
	return func() T {
		return fn(some)
	}
}
