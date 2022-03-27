package util

import "math/rand"

func Rnd(n int) func() int {
	nums := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	return func() int {
		k := rand.Intn(n)
		return nums[k]
	}
}

func GenerateRandoms(n int) []int {
	rnd := Rnd(n)
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, rnd())
	}

	return result
}
