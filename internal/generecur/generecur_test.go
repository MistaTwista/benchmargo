package generecur

import "testing"

var res int

func BenchmarkNest(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "just 10", Nums: 10},
		{Name: "just 100", Nums: 100},
		{Name: "just 1000", Nums: 1000},
	}

	for _, c := range cases {
		var r int
		//generated := generator[int](10)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = Nest(Box{i}, c.Nums)
			}
			res = r
		})
	}
}

func BenchmarkGNest(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "just generic 10", Nums: 10},
		{Name: "just generic 100", Nums: 100},
		{Name: "just generic 1000", Nums: 1000},
	}

	for _, c := range cases {
		var r int
		//generated := generator[int](10)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = GNest(GBox[int]{i}, c.Nums)
			}
			res = r
		})
	}
}
