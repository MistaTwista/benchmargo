//go:build generics

package geneuniq

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MistaTwista/benchmargo/internal/util"
)

func TestUniqG(t *testing.T) {
	cases := []struct {
		Name   string
		List   []string
		Result []string
	}{
		{
			Name:   "simple",
			List:   util.Repeat(10, []string{"one", "two", "three", "four"}...),
			Result: []string{"one", "two", "three", "four"},
		},
	}

	for _, c := range cases {
		res := GeneUniq(c.List)
		assert.Equal(t, c.Result, res)
	}
}

func TestUniqIntG(t *testing.T) {
	cases := []struct {
		Name   string
		List   []int
		Result []int
	}{
		{
			Name:   "simple",
			List:   util.Repeat(10, []int{1, 2, 3, 3, 3, 4, 5, 6, 6}...),
			Result: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, c := range cases {
		res := GeneUniq(c.List)
		assert.Equal(t, c.Result, res)
	}
}

func BenchmarkUniqStringG(b *testing.B) {
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
		list := util.Repeat(c.Nums, []string{"one", "two", "three", "four", "five"}...)
		var result []string
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = GeneUniq(list)
			}
		})
		resGlobalStr = result
	}
}

func BenchmarkUniqIntG(b *testing.B) {
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
		list := util.Repeat(c.Nums, []int{1, 2, 3, 4, 5}...)
		var result []int
		b.Run(c.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				result = GeneUniq(list)
			}
		})

		resGlobalInt = result
	}
}

type Somer interface {
	Some() error
}

type Gaga struct {
}

func (g Gaga) Some() error {
	return nil
}

func (g Gaga) String() string {
	return "hello there"
}

func Itera(list []Somer) (string, error) {
	//pee:
	for _, i := range list {
		something, ok := i.(Gaga)
		if !ok {
			return "", errors.New("bad luck")
		}
		//goto gee

		return something.String(), nil
	}

	//gee:
	//	goto pee

	return "", errors.New("nothing")
}

func TestGeneUniq(t *testing.T) {
	list := make([]Somer, 0)
	m, err := Itera(append(list, Gaga{}))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(m)
}
