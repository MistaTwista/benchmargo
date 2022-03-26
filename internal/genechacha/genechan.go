package genechan

func Processor(nums []int) []int {
	tasks := make(chan int, len(nums))
	res := make(chan int, len(nums))
	result := make([]int, 0, len(nums))

	for _, n := range nums {
		tasks <- n
	}
	close(tasks)
	for n := range tasks {
		res <- n * n
	}
	close(res)
	for d := range res {
		result = append(result, d)
	}

	return result
}

type numbers interface {
	~int | ~uint | ~float32 | ~float64
}

func GeneProcessor[T numbers](nums []T) []T {
	tasks := make(chan T, len(nums))
	res := make(chan T, len(nums))
	result := make([]T, 0, len(nums))

	for _, n := range nums {
		tasks <- n
	}
	close(tasks)
	for n := range tasks {
		res <- n * n
	}
	close(res)
	for d := range res {
		result = append(result, d)
	}

	return result
}
