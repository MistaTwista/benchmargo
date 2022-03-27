//go:build generics

package genemap

func Map[T any](list []T, cb func(T) T) []T {
	for i, itm := range list {
		list[i] = cb(itm)
	}

	return list
}

func MapAny[T any, R func(T) T](list []T, cb R) []T {
	for i, itm := range list {
		list[i] = cb(itm)
	}

	return list
}
