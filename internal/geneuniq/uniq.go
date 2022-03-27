package geneuniq

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
