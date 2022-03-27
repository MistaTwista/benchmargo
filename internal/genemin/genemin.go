package genemin

import (
	"github.com/MistaTwista/generigo/internal/util"
	"golang.org/x/exp/constraints"
)

func FindMin(list []int) int {
	var res int = 0

	for i, n := range list {
		if i == 0 {
			res = n
			continue
		}

		res = MinInt(res, n)
	}

	return res
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func GFindMin[T util.Numbers](list []T) T {
	var res T = 0

	for i, n := range list {
		if i == 0 {
			res = n
			continue
		}

		res = GMin[T](res, n)
	}

	return res
}

func GMin[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}
