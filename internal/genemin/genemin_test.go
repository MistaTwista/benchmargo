package genemin

import (
	"github.com/MistaTwista/generigo/internal/util"
	"testing"
)

var res int

func BenchmarkMin(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
		{Name: "100", Nums: 100},
		{Name: "1000", Nums: 1000},
	}

	for _, c := range cases {
		rnd := util.GgenerateRandoms(c.Nums)
		var r int
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = FindMin(rnd)
			}

			res = r
		})
	}
}

func BenchmarkGMin(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
		{Name: "100", Nums: 100},
		{Name: "1000", Nums: 1000},
	}

	for _, c := range cases {
		rnd := util.GgenerateRandoms(c.Nums)
		var r int
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GFindMin(rnd)
			}

			res = r
		})
	}
}
