//go:build generics

package genechan

import (
	"github.com/MistaTwista/generigo/internal/util"
	"testing"
)

func TestGeneProcessor(t *testing.T) {
	x := GeneProcessor(util.GgenerateRandoms[int](10), 10, 10)
	t.Log(x, len(x))
}

var res []int

func BenchmarkGeneProcessor(b *testing.B) {
	b.Skip()
	cases := []struct {
		Name    string
		Nums    int
		Workers int
		Repeats int
	}{
		{Name: "10-2-1", Nums: 10, Workers: 2, Repeats: 1},
	}

	for _, c := range cases {
		var r []int
		nums := util.GgenerateRandoms(c.Nums)

		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GeneProcessor(nums, c.Workers, c.Repeats)
			}
			res = r
		})
	}
}
