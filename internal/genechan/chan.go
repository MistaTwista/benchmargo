package genechan

import (
	"sync"
)

func Processor(nums []int, workers int, repeats int) []int {
	tasks := make(chan int, 10)
	res := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 0; i <= workers; i++ {
		wg.Add(1)
		go func(taskChan chan int) {
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

	result := make([]int, 0, len(nums))
	go func(resultsChan chan int) {
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
