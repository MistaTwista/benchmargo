package geneslic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	cases := []struct {
		Name   string
		List   []string
		Result []string
	}{
		{
			Name:   "simple",
			List:   generateFakesStr(10),
			Result: []string{"one", "two", "three", "four"},
		},
	}

	for _, c := range cases {
		res := UniqStrings(c.List)
		assert.Equal(t, c.Result, res)
	}
}

func TestGeneuniq(t *testing.T) {
	cases := []struct {
		Name   string
		List   []string
		Result []string
	}{
		{
			Name:   "simple",
			List:   generateFakesStr(10),
			Result: []string{"one", "two", "three", "four"},
		},
	}

	for _, c := range cases {
		res := GeneUniq(c.List)
		assert.Equal(t, c.Result, res)
	}
}

func generateFakesStr(num int) []string {
	fromList := []string{"one", "two", "three", "four"}
	result := make([]string, 0, num*len(fromList))
	for _, s := range fromList {
		result = append(result, s)
	}

	return result
}

func generateFakesInt(num int) []int {
	fromList := []int{1, 2, 3, 4}
	result := make([]int, 0, num*len(fromList))
	for _, s := range fromList {
		result = append(result, s)
	}

	return result
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
		list := generateFakesStr(c.Nums)
		var result []string
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = UniqStrings(list)
			}
		})
		resGlobalStr = result
	}
}

func BenchmarkGeneUniq(b *testing.B) {
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
		list := generateFakesStr(c.Nums)
		var result []string
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = GeneUniq(list)
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
		list := generateFakesInt(c.Nums)
		var result []int
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = UniqInts(list)
			}
		})
		resGlobalInt = result
	}
}

func BenchmarkGeneUniqInt(b *testing.B) {
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
		list := generateFakesInt(c.Nums)
		var result []int
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = GeneUniq(list)
			}
		})
		resGlobalInt = result
	}
}
