package genestat

import "testing"

var rG uint64
var wG uint64

func BenchmarkJustProcessor(b *testing.B) {
	cases := []struct {
		Name   string
		RCount int
		WCount int
	}{
		{Name: "just-100-10", RCount: 100, WCount: 10},
		{Name: "just-1000-100", RCount: 1000, WCount: 100},
		{Name: "just-10000-1000", RCount: 10000, WCount: 1000},
	}

	for _, c := range cases {
		var r uint64
		var w uint64
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r, w = Processor(c.RCount, c.WCount)
			}
			rG = r
			wG = w
		})

		b.Log(rG, wG)
	}
}
