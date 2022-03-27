//go:build generics

package geneuniq

func GeneUniq[T comparable](list []T) []T {
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
