package genestack

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() interface{} {
	x := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return x
}

type GStack[T any] struct {
	values []T
}

func (s *GStack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *GStack[T]) Pop() T {
	x := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return x
}
