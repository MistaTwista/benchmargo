package genechan

import (
	"sync"
)

func Slice2Chan(repeats int, nums ...int) <-chan int {
	res := make(chan int, 10)
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

func Chan2Slice(resultsChan <-chan int) []int {
	res := make([]int, 0, 10)
	for r := range resultsChan {
		res = append(res, r)
	}

	return res
}

func Worker(tasksChan <-chan int) <-chan int {
	res := make(chan int)

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

func Mix(chList []<-chan int) <-chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range chList {
		wg.Add(1)
		go func(dataCh <-chan int) {
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

func Processor(nums []int, workers int, repeats int) []int {
	tasks := Slice2Chan(repeats, nums...)
	resChans := make([]<-chan int, 0, workers)
	for i := 0; i <= workers; i++ {
		resChans = append(resChans, Worker(tasks))
	}

	return Chan2Slice(Mix(resChans))
}
