//go:build generics

package genelast

import (
	"sync"
)

func gPutInChan[T any](ch chan T, some T) {
	ch <- some
}

func GLast[T any](items ...T) T {
	var result T
	resCh := make(chan T)
	defer func() {
		close(resCh)
	}()

	var wg sync.WaitGroup
	go func() {
		for n := range resCh {
			wg.Done()
			result = n
		}
	}()

	for _, itm := range items {
		wg.Add(1)
		go gPutInChan(resCh, itm)
	}

	wg.Wait()
	return result
}

func LastWithG(nums ...int) int {
	var result int
	resCh := make(chan int)
	defer func() {
		close(resCh)
	}()

	var wg sync.WaitGroup
	go func() {
		for n := range resCh {
			wg.Done()
			result = n
		}
	}()

	for _, n := range nums {
		wg.Add(1)
		go gPutInChan(resCh, n)
	}

	wg.Wait()
	return result
}
