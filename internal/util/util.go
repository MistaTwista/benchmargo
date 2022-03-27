package util

import "math/rand"

type Stringer interface {
	String() string
}

func StringsToInterface(list ...string) []interface{} {
	interfacedList := make([]interface{}, 0, len(list))
	for _, itm := range list {
		interfacedList = append(interfacedList, itm)
	}

	return interfacedList
}
func FloatsToInterface(list ...float64) []interface{} {
	interfacedList := make([]interface{}, 0, len(list))
	for _, itm := range list {
		interfacedList = append(interfacedList, itm)
	}

	return interfacedList
}

func IntsToInterface(list ...int) []interface{} {
	interfacedList := make([]interface{}, 0, len(list))
	for _, itm := range list {
		interfacedList = append(interfacedList, itm)
	}

	return interfacedList
}

func RepeatFloats(num int, seed ...float64) []float64 {
	res := make([]float64, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, seed...)
	}

	return res
}

func RepeatInts(num int, seed ...int) []int {
	res := make([]int, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, seed...)
	}

	return res
}

func RepeatStrings(num int, seed ...string) []string {
	res := make([]string, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, seed...)
	}

	return res
}

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
