package genestat

import "testing"

var rG uint64
var wG uint64

func BenchmarkProcessor(b *testing.B) {
	cases := []struct {
		Name    string
		Nums    int
		Workers int
		Repeats int
	}{
		{Name: "10-2-1", Nums: 10, Workers: 2, Repeats: 1},
	}

	for _, c := range cases {
		var r uint64
		var w uint64
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r, w = Processor()
			}
			rG = r
			wG = w
		})

		b.Log(rG, wG)
	}
}

func BenchmarkGProcessor(b *testing.B) {
	cases := []struct {
		Name    string
		Nums    int
		Workers int
		Repeats int
	}{
		{Name: "10-2-1", Nums: 10, Workers: 2, Repeats: 1},
	}

	for _, c := range cases {
		var r uint64
		var w uint64
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r, w = GProcessor[float64]()
			}
			rG = r
			wG = w
		})

		b.Log(rG, wG)
	}
}
