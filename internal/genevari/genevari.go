package genevari

import "github.com/MistaTwista/generigo/internal/util"

func Pipe(fns ...func(int) int) int {
	var r = int(0)
	for _, fn := range fns {
		r = fn(r)
	}

	return r
}

func Identity(n int) func(int) int {
	return func(_ int) int {
		return n
	}
}

var Multiply = func(i int) int {
	return i * i
}

var Plus = func(i int) int {
	return i + i
}

func MultiplyWith(n int) func(int) int {
	return func(i int) int {
		return i * n
	}
}
func GPipe[T util.Numbers](fns ...func(T) T) T {
	var r = T(0)
	for _, fn := range fns {
		r = fn(r)
	}

	return r
}

func GIdentity[T util.Numbers](n T) func(T) T {
	return func(_ T) T {
		return n
	}
}

func GMultiply[T util.Numbers](i T) T {
	return i * i
}

func GPlus[T util.Numbers](i T) T {
	return i + i
}

func GMultiplyWith[T util.Numbers](n T) func(T) T {
	return func(i T) T {
		return i * n
	}
}
