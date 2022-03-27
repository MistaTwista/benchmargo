package genechan

import (
	"github.com/MistaTwista/generigo/internal/util"
	"testing"
)

func TestJustProcessor(t *testing.T) {
	x := Processor([]int{1, 2, 3}, 10, 10)
	t.Log(x)
}

func BenchmarkJustProcessor(b *testing.B) {
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
		nums := util.GenerateRandoms(c.Nums)
		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Processor(nums, c.Workers, c.Repeats)
			}
			res = r
		})
	}
}
