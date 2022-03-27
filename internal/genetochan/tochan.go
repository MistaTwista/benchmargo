package genetochan

import "sync"

func ToChan(fn func() int) <-chan int {
	res := make(chan int)
	go func() {
		defer func() {
			close(res)
		}()

		x := fn()
		res <- x
	}()

	return res
}

func Mix(chans ...<-chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)
		go func(dataCh <-chan int) {
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
