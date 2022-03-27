//go:build generics

package genecon

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func BenchmarkConG(b *testing.B) {
	cases := []struct {
		Name  string
		Items []Str
	}{
		{Name: "10", Items: gena(10)},
		{Name: "100", Items: gena(100)},
	}

	for _, c := range cases {
		var r []string
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GStringify(c.Items...)
			}
			res = r
		})
	}
}

func BenchmarkConInterfaceG(b *testing.B) {
	cases := []struct {
		Name  string
		Items []util.Stringer
	}{
		{Name: "10", Items: NewStringer(gena(10))},
		{Name: "100", Items: NewStringer(gena(100))},
	}

	for _, c := range cases {
		var r []string
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = GStringify(c.Items...)
			}
			res = r
		})
	}
}
