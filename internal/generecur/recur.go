package generecur

type Box struct {
	value int
}

func Nest(b Box, n int) int {
	if n == 0 {
		return b.value
	}

	return Nest(Box{b.value + 1}, n-1)
}
