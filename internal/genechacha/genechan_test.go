package genechan

import "testing"

var res []int

func generator[T numbers](num int) []T {
	res := make([]T, 0, num)
	for i := 0; i <= num; i++ {
		res = append(res, T(i))
	}

	return res
}

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
		var r []int
		//generated := generator[int](10)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = Processor([]int{1, 2, 3})
			}
			res = r
		})
		b.Log(res)
	}
}

func BenchmarkGeneProcessor(b *testing.B) {
	cases := []struct {
		Name    string
		Nums    int
		Workers int
		Repeats int
	}{
		{Name: "10-2-1", Nums: 10, Workers: 2, Repeats: 1},
	}

	for _, c := range cases {
		var r []int
		//generated := generator[int](1000)
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				//b.Log(i, b.N)
				r = GeneProcessor([]int{1, 2, 3})
			}
			res = r
		})
		b.Log(res)
	}
}
