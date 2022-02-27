package genechan

func Proc(num int) int {
	resCh := make(chan int)

	go func(some int) {
		resCh <- some
	}(num)

	return <-resCh
}

type numbers interface {
	~int | ~uint | ~float32 | ~float64
}

func gchanputin[T numbers](some T, ch chan T) {
	ch <- some
}

func GProc[T numbers](n T) T {
	resCh := make(chan T)
	go gchanputin(n, resCh)

	return <-resCh
}

//func GeneProcessor[T numbers](nums []T, workers int, repeats int) []T {
//	tasks := make(chan T, workers)
//	res := make(chan T, workers)
//
//	var wg sync.WaitGroup
//	for i := 0; i <= workers; i++ {
//		wg.Add(1)
//		go func(taskChan chan T) {
//			defer wg.Done()
//			for {
//				select {
//				case t, ok := <-taskChan:
//					if !ok {
//						return
//					}
//
//					res <- t * t
//				}
//			}
//		}(tasks)
//	}
//
//	result := make([]T, 0, len(nums))
//	go func(resultsChan chan T) {
//		for {
//			select {
//			case r := <-resultsChan:
//				result = append(result, r)
//			}
//		}
//	}(res)
//
//	for _, n := range nums {
//		for i := 0; i < repeats; i++ {
//			tasks <- n
//		}
//	}
//	close(tasks)
//
//	wg.Wait()
//	close(res)
//
//	return result
//}