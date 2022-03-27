package genechacha

import (
	"github.com/MistaTwista/generigo/internal/util"
	"testing"
)

var res []int

func BenchmarkJustProcessor(b *testing.B) {
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
		var r []int
		nums := util.GenerateRandoms(c.Nums)
		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Processor(nums)
			}
			res = r
		})
	}
}
