package generecur

type Box struct {
	value int
}

func sum(b Box, n int) int {
	if n == 0 {
		return b.value
	}

	return sum(Box{b.value + 1}, n-1)
}

func Sum(a, b int) int {
	return sum(Box{a}, b)
}
