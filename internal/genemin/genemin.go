//go:build generics

package genemin

import (
	"golang.org/x/exp/constraints"

	"github.com/MistaTwista/benchmargo/internal/util"
)

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

func GFindMinEmbed[T util.Numbers](list []T) T {
	var res T = 0

	for i, n := range list {
		if i == 0 {
			res = n
			continue
		}

		if res > n {
			res = n
		}
	}

	return res
}
