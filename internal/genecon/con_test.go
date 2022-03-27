package genecon

import (
	"testing"

	"github.com/MistaTwista/generigo/internal/util"
)

func BenchmarkJustStringify(b *testing.B) {
	cases := []struct {
		Name  string
		Items []Str
	}{
		{Name: "10", Items: gena(10)},
		{Name: "100", Items: gena(100)},
	}

	for _, c := range cases {
		var r []string

		list := make([]util.Stringer, 0, len(c.Items))
		for _, itm := range c.Items {
			list = append(list, itm)
		}

		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Stringify(list)
			}
			res = r
		})
	}
}

func BenchmarkJustStringifyInterface(b *testing.B) {
	cases := []struct {
		Name  string
		Items []util.Stringer
	}{
		{Name: "10", Items: NewStringer(gena(10))},
		{Name: "100", Items: NewStringer(gena(100))},
	}

	for _, c := range cases {
		var r []string
		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Stringify(c.Items)
			}
			res = r
		})
	}
}
