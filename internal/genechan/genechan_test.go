//go:build generics

package genechan

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func TestChanG(t *testing.T) {
	list := GeneProcessor([]int{1, 2, 3}, 10, 1)
	var sum int
	for _, x := range list {
		sum += x
	}
	res := 14
	if sum != res {
		t.Fatalf("%d != %d", sum, res)
	}
}

func BenchmarkChanG(b *testing.B) {
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
		nums := util.GenerateRandoms(c.Nums)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GeneProcessor(nums, c.Workers, c.Repeats)
			}
			res = r
		})
	}
}
