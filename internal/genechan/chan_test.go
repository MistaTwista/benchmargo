package genechan

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func TestChanJ(t *testing.T) {
	//t.Skip()
	list := Processor([]int{1, 2, 3}, 10, 1)
	var sum int
	for _, x := range list {
		sum += x
	}
	t.Log(sum)
	res := 14
	if sum != res {
		t.Fatalf("%d != %d", sum, res)
	}
}

var res []int

func BenchmarkChanJ(b *testing.B) {
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
				r = Processor(nums, c.Workers, c.Repeats)
			}
			res = r
		})
	}
}
