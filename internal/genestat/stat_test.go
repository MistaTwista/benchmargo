package genestat

import "testing"

var rG uint64
var wG uint64

func BenchmarkStatJ(b *testing.B) {
	cases := []struct {
		Name   string
		RCount int
		WCount int
	}{
		{Name: "100-10", RCount: 100, WCount: 10},
		{Name: "10-100", RCount: 10, WCount: 100},
		{Name: "1000-100", RCount: 1000, WCount: 100},
		{Name: "10000-1000", RCount: 10000, WCount: 1000},
		{Name: "10000-10000", RCount: 10000, WCount: 10000},
		{Name: "100000-100000", RCount: 100000, WCount: 100000},
		{Name: "1000000-1000000", RCount: 1000000, WCount: 1000000},
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
