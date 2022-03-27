//go:build generics

package genemap

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func BenchmarkMapG(b *testing.B) {
	cases := []struct {
		Name  string
		Items []int
	}{
		{Name: "3", Items: []int{1, 2, 3}},
		{Name: "10", Items: util.GenerateRandoms(10)},
		{Name: "100", Items: util.GenerateRandoms(100)},
		{Name: "1000", Items: util.GenerateRandoms(1000)},
		{Name: "10000", Items: util.GenerateRandoms(10000)},
	}

	for _, c := range cases {
		var r []int
		b.Run(c.Name+" typed arg", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = Map(c.Items, func(i int) int {
					return i * 2
				})
			}
		})
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = MapAny(c.Items, func(i int) int {
					return i * 2
				})
			}
		})
		result = r
	}
}
