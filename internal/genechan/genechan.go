//go:build generics

package genechan

import (
	"github.com/MistaTwista/generigo/internal/util"
	"sync"
)

func GeneProcessor[T util.Numbers](nums []T, workers int, repeats int) []T {
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
