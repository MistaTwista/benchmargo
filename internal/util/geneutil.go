//go:build generics

package util

import "math/rand"

type Numbers interface {
	~int | ~uint | ~float32 | ~float64
}

type Addeable interface {
	~int | ~float64 | ~string
}

type Checkable interface {
	~int | ~string | ~float32 | ~float64
}

func Repeat[T any](times int, seed ...T) []T {
	res := make([]T, 0, times)
	for i := 0; i < times; i++ {
		res = append(res, seed...)
	}

	return res
}

func ToInterface[T any](list ...T) []interface{} {
	interfacedList := make([]interface{}, 0, len(list))
	for _, itm := range list {
		interfacedList = append(interfacedList, itm)
	}

	return interfacedList
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
