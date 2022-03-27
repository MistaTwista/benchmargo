package genestack

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsGotSome() bool {
	return len(s.values) != 0
}

func (s *Stack) Pop() interface{} {
	x := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return x
}
