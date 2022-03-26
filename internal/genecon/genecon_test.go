package genecon

import (
	"strconv"
	"testing"
)

type Str struct {
	S string
}

func (s Str) String() string {
	return s.S
}

var res []string

func gena(n int) []Str {
	res := make([]Str, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, Str{S: strconv.Itoa(i)})
	}
	return res
}

func NewStringer(list []Str) []Stringer {
	res := make([]Stringer, 0, len(list))
	for _, itm := range list {
		res = append(res, Stringer(itm))
	}

	return res
}

func BenchmarkGProcessor(b *testing.B) {
	cases := []struct {
		Name  string
		Items []Str
	}{
		{Name: "just generic 10", Items: gena(10)},
		{Name: "just generic 100", Items: gena(100)},
	}

	for _, c := range cases {
		var r []string
		//generated := generator[int](10)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = GStringify(c.Items)
			}
			res = r
		})
		//b.Log(res)
	}
}

func BenchmarkGProcessorInterface(b *testing.B) {
	cases := []struct {
		Name  string
		Items []Stringer
	}{
		{Name: "generic interface iteration 10", Items: NewStringer(gena(10))},
		{Name: "generic interface iteration 100", Items: NewStringer(gena(100))},
	}

	for _, c := range cases {
		var r []string
		//generated := generator[int](10)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = GStringify(c.Items)
			}
			res = r
		})
		//b.Log(res)
	}
}

func BenchmarkProcessor(b *testing.B) {
	cases := []struct {
		Name  string
		Items []Str
	}{
		{Name: "non generic 10", Items: gena(10)},
		{Name: "non generic 100", Items: gena(100)},
	}

	for _, c := range cases {
		var r []string
		//generated := generator[int](10)
		list := make([]Stringer, 0, len(c.Items))
		for _, itm := range c.Items {
			list = append(list, itm)
		}

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = Stringify(list)
			}
			res = r
		})
		//b.Log(res)
	}
}

func BenchmarkProcessorInterface(b *testing.B) {
	cases := []struct {
		Name  string
		Items []Stringer
	}{
		{Name: "non generic by interface 10", Items: NewStringer(gena(10))},
		{Name: "non generic by interface 100", Items: NewStringer(gena(100))},
	}

	for _, c := range cases {
		var r []string
		//generated := generator[int](10)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = Stringify(c.Items)
			}
			res = r
		})
		//b.Log(res)
	}
}
