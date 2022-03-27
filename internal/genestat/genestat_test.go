//go:build generics

package genestat

import "testing"

func BenchmarkGeneProcessor(b *testing.B) {
	cases := []struct {
		Name   string
		RCount int
		WCount int
	}{
		{Name: "generics-100-10", RCount: 100, WCount: 10},
		{Name: "generics-1000-100", RCount: 1000, WCount: 100},
		{Name: "generics-10000-1000", RCount: 10000, WCount: 1000},
	}

	for _, c := range cases {
		var r uint64
		var w uint64
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r, w = GProcessor[int](c.RCount, c.WCount)
			}
			rG = r
			wG = w
		})

		b.Log(rG, wG)
	}
}
