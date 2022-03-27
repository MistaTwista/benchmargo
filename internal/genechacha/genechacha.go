//go:build generics

package genechacha

import "github.com/MistaTwista/generigo/internal/util"

func GeneProcessor[T util.Numbers](nums []T) []T {
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
