package genechan

import (
	"github.com/MistaTwista/generigo/internal/util"
	"testing"
)

func TestAsync(t *testing.T) {
	x := Processor([]int{1, 2, 3}, 10, 10)
	t.Log(x)
}

func TestGeneAsync(t *testing.T) {
	x := GeneProcessor(util.GgenerateRandoms[int](10), 10, 10)
	t.Log(x, len(x))
}

var res []int

func BenchmarkProcessor(b *testing.B) {
	//b.Skip()
	cases := []struct {
		Name    string
		Nums    int
		Workers int
		Repeats int
	}{
		{Name: "10-2-1", Nums: 10, Workers: 2, Repeats: 1},
		//{Name: "100-2-1", Nums: 100, Workers: 2, Repeats: 1},
		//{Name: "100-4-1", Nums: 100, Workers: 4, Repeats: 1},
		//{Name: "100-8-1", Nums: 100, Workers: 8, Repeats: 1},
		//{Name: "100-16-1", Nums: 100, Workers: 16, Repeats: 1},
		//{Name: "100-32-1", Nums: 100, Workers: 32, Repeats: 1},
		//{Name: "1000-1-1", Nums: 1000, Workers: 1, Repeats: 1},
		//{Name: "1000-2-1", Nums: 1000, Workers: 2, Repeats: 1},
		//{Name: "1000-4-1", Nums: 1000, Workers: 4, Repeats: 1},
		//{Name: "1000-8-1", Nums: 1000, Workers: 8, Repeats: 1},
		//{Name: "1000-16-1", Nums: 1000, Workers: 16, Repeats: 1},
		//{Name: "1000-32-1", Nums: 1000, Workers: 32, Repeats: 1},
	}

	for _, c := range cases {
		var r []int
		nums := util.GgenerateRandoms(c.Nums)
		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Processor(nums, c.Workers, c.Repeats)
			}
			res = r
		})

		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GeneProcessor(nums, c.Workers, c.Repeats)
			}
			res = r
		})
	}
}
