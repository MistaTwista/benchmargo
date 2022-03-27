package geneslic

import (
	"github.com/MistaTwista/generigo/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJustUniq(t *testing.T) {
	cases := []struct {
		Name   string
		List   []string
		Result []string
	}{
		{
			Name:   "simple",
			List:   util.GenAddeableStrings(10, []string{"one", "two", "three", "four"}...),
			Result: []string{"one", "two", "three", "four"},
		},
	}

	for _, c := range cases {
		res := UniqStrings(c.List)
		assert.Equal(t, c.Result, res)
	}
}

var resGlobalStr []string

func BenchmarkUniqString(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
		{Name: "100", Nums: 100},
		{Name: "1000", Nums: 1000},
		{Name: "10000", Nums: 10000},
	}

	for _, c := range cases {
		list := util.GenAddeableStrings(c.Nums, []string{"one", "two", "three", "four", "five"}...)
		var result []string
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = UniqStrings(list)
			}
		})
		resGlobalStr = result
	}
}

var resGlobalInt []int

func BenchmarkUniqInt(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
		{Name: "100", Nums: 100},
		{Name: "1000", Nums: 1000},
		{Name: "10000", Nums: 10000},
	}

	for _, c := range cases {
		list := util.GenAddeableInts(c.Nums, []int{1, 2, 3, 4, 5}...)
		var result []int
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = UniqInts(list)
			}
		})
		resGlobalInt = result
	}
}
