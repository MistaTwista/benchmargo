//go:build generics

package geneapply

import (
	"github.com/MistaTwista/benchmargo/internal/util"
)

func GApply[T any](fns ...func(T) T) T {
	var r T
	for _, fn := range fns {
		r = fn(r)
	}

	return r
}

func GIdentity[T any](n T) func(T) T {
	return func(_ T) T {
		return n
	}
}

func GMultiply[T util.Numbers](i T) T {
	return i * i
}

func GPlus[T util.Addable](i T) T {
	return i + i
}

func GMultiplyWith[T util.Numbers](n T) func(T) T {
	return func(i T) T {
		return i * n
	}
}
