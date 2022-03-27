//go:build generics

package generecur

import "github.com/MistaTwista/benchmargo/internal/util"

type GBox[T util.Numbers] struct {
	value T
}

func gsum[T util.Numbers](b GBox[T], n int) T {
	if n == 0 {
		return b.value
	}

	return gsum(GBox[T]{b.value + 1}, n-1)
}

func GSum[T util.Numbers](a T, b int) T {
	return gsum[T](GBox[T]{a}, b)
}

//type AnyBox[A any] struct {
//	value A
//}
//
//func Nest[A any](b AnyBox[A], n int) interface{} {
//	if n == 0 {
//		return b
//	}
//
//	return Nest(AnyBox[AnyBox[A]]{b}, n-1)
//}
