package genepower

func Power(nums []int) []int {
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
