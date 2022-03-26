package genestack

import "testing"

func TestStack_Pop(t *testing.T) {
	s := Stack{}
	s.Push(42)
	s.Push(43)
	s.Push("some")
	s.Push([]string{"one", "two"})
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}

func TestStack_GPop(t *testing.T) {
	s := GStack[int]{}
	s.Push(42)
	s.Push(43)
	t.Log(s.Pop())
	t.Log(s.Pop())
}
