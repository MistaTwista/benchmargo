package generecur

import (
	"fmt"
	"testing"
)

func TestNest(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   int
		Result int
	}{
		{Name: "just 10", Nums: 10, Result: 20},
		{Name: "just 100", Nums: 100, Result: 200},
		{Name: "just 1000", Nums: 1000, Result: 2000},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			r := Nest(Box{c.Nums}, c.Nums)
			if r != c.Result {
				t.Fatal(fmt.Sprintf("want: %d, got: %d", c.Result, r))
			}
		})
	}
}

var res int

func BenchmarkNest(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
		{Name: "100", Nums: 100},
		{Name: "1000", Nums: 1000},
		{Name: "10000", Nums: 10000},
		{Name: "100000", Nums: 100000},
	}

	for _, c := range cases {
		var r int
		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Nest(Box{i}, c.Nums)
			}
			res = r
		})

		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GNest(GBox[int]{i}, c.Nums)
			}
			res = r
		})
	}
}
