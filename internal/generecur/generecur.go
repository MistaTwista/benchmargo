//go:build generics

package generecur

import "github.com/MistaTwista/generigo/internal/util"

type GBox[T util.Numbers] struct {
	value T
}

func GNest[T util.Numbers](b GBox[T], n int) T {
	if n == 0 {
		return b.value
	}

	return GNest(GBox[T]{b.value + 1}, n-1)
}
