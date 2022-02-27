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
			List:   generateFakes(10),
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
			List:   generateFakes(10),
			Result: []string{"one", "two", "three", "four"},
		},
	}

	for _, c := range cases {
		res := GeneUniq(c.List)
		assert.Equal(t, c.Result, res)
	}
}

func generateFakes(num int) []string {
	fromList := []string{"one", "two", "three", "four"}
	result := make([]string, 0, num*len(fromList))
	for _, s := range fromList {
		result = append(result, s)
	}

	return result
}

var resGlobal []string

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
		list := generateFakes(c.Nums)
		var result []string
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = UniqStrings(list)
			}
		})
		resGlobal = result
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
		list := generateFakes(c.Nums)
		var result []string
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = GeneUniq(list)
			}
		})
		resGlobal = result
	}
}
