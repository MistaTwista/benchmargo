package util

import "math/rand"

func Rnd(n int) func() int {
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, int(i))
	}

	return func() int {
		k := rand.Intn(5)
		return nums[k]
	}
}

type Numbers interface {
	~int | ~uint | ~float32 | ~float64
}

func Grnd[T Numbers](n T) func() T {
	nums := make([]T, 0, int(n))
	for i := 1; i <= int(n); i++ {
		nums = append(nums, T(i))
	}

	return func() T {
		k := rand.Intn(int(n))
		return nums[k]
	}
}

func GgenerateRandoms[T Numbers](n T) []T {
	rnd := Grnd[T](n)
	result := make([]T, 0, int(n))
	for i := 0; i < int(n); i++ {
		result = append(result, rnd())
	}

	return result
}
