package genepower

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowerJ(t *testing.T) {
	cases := []struct {
		Name   string
		Nums   []int
		Result []int
	}{
		{Name: "just", Nums: []int{1, 2, 3, 4}, Result: []int{1, 4, 9, 16}},
	}

	for _, c := range cases {
		r := Power(c.Nums)
		assert.Equal(t, r, c.Result)
	}
}

var res []int

func BenchmarkPowerJ(b *testing.B) {
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
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Power(nums)
			}
			res = r
		})
	}
}
