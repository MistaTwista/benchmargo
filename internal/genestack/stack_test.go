package genestack

import (
	"fmt"
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func TestStackJ(t *testing.T) {
	s := Stack{}
	s.Push(42)
	s.Push(43)
	s.Push("some")
	s.Push([]string{"one", "two"})
	for _, item := range []interface{}{[]string{"one", "two"}, "some", 43, 42} {
		cur, ok := s.Pop()
		if !ok {
			t.Fatal(fmt.Sprintf("nothing left, but you asked %v", item))
		}

		want := fmt.Sprintf("%v", item)
		got := fmt.Sprintf("%v", cur)

		if want != got {
			t.Fatal(fmt.Sprintf("want: %s, got: %s", want, got))
		}
	}
}

var dataInterface interface{}

func BenchmarkStackJ(b *testing.B) {
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

				for {
					cur, ok := s.Pop()
					if !ok {
						break
					}
					result = cur
				}
			}
		})
		dataInterface = result
	}
}
