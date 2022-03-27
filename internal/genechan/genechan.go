//go:build generics

package genechan

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"sync"
)

func GSlice2Chan[T any](repeats int, nums ...T) <-chan T {
	res := make(chan T, 10)
	go func() {
		defer func() {
			close(res)
		}()

		for _, n := range nums {
			for i := 0; i < repeats; i++ {
				res <- n
			}
		}
	}()

	return res
}

func GChan2Slice[T any](resultsChan <-chan T) []T {
	res := make([]T, 0, 10)
	for r := range resultsChan {
		res = append(res, r)
	}

	return res
}

func GWorker[T util.Numbers](tasksChan <-chan T) <-chan T {
	res := make(chan T)

	go func() {
		defer func() {
			close(res)
		}()

		for t := range tasksChan {
			res <- t * t
		}
	}()

	return res
}

func GMix[T any](chList []<-chan T) <-chan T {
	res := make(chan T)
	var wg sync.WaitGroup
	for _, ch := range chList {
		wg.Add(1)
		go func(dataCh <-chan T) {
			defer wg.Done()
			for d := range dataCh {
				res <- d
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func GeneProcessor[T util.Numbers](nums []T, workers int, repeats int) []T {
	tasks := GSlice2Chan[T](repeats, nums...)
	resChans := make([]<-chan T, 0, workers)
	for i := 0; i <= workers; i++ {
		resChans = append(resChans, GWorker(tasks))
	}

	return GChan2Slice(GMix(resChans))
}

func GenePreProcessor[T util.Numbers](nums []T, workers int, repeats int) []T {
	tasks := make(chan T)
	res := make(chan T)

	var wg sync.WaitGroup
	for i := 0; i <= workers; i++ {
		wg.Add(1)
		go func(taskChan chan T) {
			defer wg.Done()
			for {
				select {
				case t, ok := <-taskChan:
					if !ok {
						return
					}

					res <- t * t
				}
			}
		}(tasks)
	}

	result := make([]T, 0, len(nums))
	go func(resultsChan chan T) {
		for {
			select {
			case r := <-resultsChan:
				result = append(result, r)
			}
		}
	}(res)

	for _, n := range nums {
		for i := 0; i < repeats; i++ {
			tasks <- n
		}
	}
	close(tasks)

	wg.Wait()
	close(res)

	return result
}
