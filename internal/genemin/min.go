package genemin

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
