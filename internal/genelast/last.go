package genelast

import "sync"

func putInChan(ch chan int, some int) {
	ch <- some
}

func Last(nums ...int) int {
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
		go putInChan(resCh, n)
	}

	wg.Wait()
	return result
}

func putAnyInChan(ch chan interface{}, n interface{}) {
	ch <- n
}

func LastInterface(nums ...interface{}) interface{} {
	var result interface{}
	resCh := make(chan interface{})
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
		go putAnyInChan(resCh, n)
	}

	wg.Wait()
	return result
}
