package geneslic

func UniqInts(list []int) []int {
	uniq := make([]int, 0, len(list))
	index := make(map[int]struct{})

	for _, s := range list {
		if _, ok := index[s]; ok {
			continue
		}

		index[s] = struct{}{}
		uniq = append(uniq, s)
	}

	return uniq
}

func UniqStrings(list []string) []string {
	uniq := make([]string, 0, len(list))
	index := make(map[string]struct{})

	for _, s := range list {
		if _, ok := index[s]; ok {
			continue
		}

		index[s] = struct{}{}
		uniq = append(uniq, s)
	}

	return uniq
}

type checkable interface {
	~int | ~string | ~float32 | ~float64
}

func GeneUniq[T checkable](list []T) []T {
	uniq := make([]T, 0, len(list))
	index := make(map[T]struct{})

	for _, itm := range list {
		if _, ok := index[itm]; ok {
			continue
		}

		index[itm] = struct{}{}
		uniq = append(uniq, itm)
	}

	return uniq
}
