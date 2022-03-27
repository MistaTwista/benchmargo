//go:build generics

package genestack

type GStack[T any] struct {
	values []T
}

func (s *GStack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *GStack[T]) Pop() (T, bool) {
	var x T
	if len(s.values) == 0 {
		return x, false
	}

	x = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return x, true
}
