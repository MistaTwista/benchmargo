package genestack

import (
	"fmt"
	"github.com/MistaTwista/generigo/internal/util"
	"testing"
)

func TestJustStack(t *testing.T) {
	s := Stack{}
	s.Push(42)
	s.Push(43)
	s.Push("some")
	s.Push([]string{"one", "two"})
	for _, item := range []interface{}{[]string{"one", "two"}, "some", 43, 42} {
		if ok := s.IsGotSome(); !ok {
			t.Fatal(fmt.Sprintf("nothing left, but you asked %v", item))
		}

		want := fmt.Sprintf("%v", item)
		got := fmt.Sprintf("%v", s.Pop())

		if want != got {
			t.Fatal(fmt.Sprintf("want: %s, got: %s", want, got))
		}
	}
}

var dataInterface interface{}

func BenchmarkJustStack(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
		{Name: "100", Nums: 100},
		{Name: "1000", Nums: 1000},
		{Name: "10000", Nums: 10000},
	}

	for _, c := range cases {
		list := util.RepeatInts(c.Nums, []int{1, 2, 3, 4, 5}...)
		var result interface{}
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				s := Stack{}
				for _, x := range list {
					s.Push(x)
				}

				for s.IsGotSome() {
					result = s.Pop()
				}
			}
		})
		dataInterface = result
	}
}
