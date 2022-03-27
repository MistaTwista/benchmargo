//go:build generics

package genemin

import (
	"fmt"
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func TestMinG(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Result int
	}{
		{Name: "just", Nums: []int{9, 8, 7, 6, 5, 1, 4}, Result: 1},
	}

	for _, c := range cases {
		t.Run(c.Name+"-generics", func(t *testing.T) {
			r := GFindMin(c.Nums)
			if r != c.Result {
				t.Fatal(fmt.Sprintf("want: %d, got: %d", c.Result, r))
			}
		})
	}
}

func BenchmarkMinG(b *testing.B) {
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
		rnd := util.GgenerateRandoms(c.Nums)
		var r int
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GFindMin(rnd)
			}

			res = r
		})

		b.Run(c.Name+" embed", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GFindMinEmbed(rnd)
			}

			res = r
		})
	}
}
