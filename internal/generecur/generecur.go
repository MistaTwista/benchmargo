package generecur

type GBox[A int | int64 | int32] struct {
	value A
}

func GNest[A int | int64 | int32](b GBox[A], n int) A {
	if n == 0 {
		return b.value
	}

	return GNest(GBox[A]{b.value}, n-1)
}

type Box struct {
	value int
}

func Nest(b Box, n int) int {
	if n == 0 {
		return b.value
	}

	return Nest(Box{b.value}, n-1)
}
