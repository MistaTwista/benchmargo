package genechan

import "testing"

func TestProc(t *testing.T) {
	r := Proc(42)
	if r != 42 {
		t.Fail()
	}
}

func TestGProc(t *testing.T) {
	r := GProc(42)
	if r != 42 {
		t.Fail()
	}
}

var res int

func BenchmarkProcessor(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = Proc(42)
	}
	res = r
}

func BenchmarkGProcessor(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = GProc(42)
	}
	res = r
}

//func BenchmarkGeneProcessor(b *testing.B) {
//	cases := []struct {
//		Name    string
//		Nums    int
//		Workers int
//		Repeats int
//	}{
//		{Name: "10-2-1", Nums: 10, Workers: 2, Repeats: 1},
//		//{Name: "100-2-1", Nums: 100, Workers: 2, Repeats: 1},
//		{Name: "100-4-1", Nums: 100, Workers: 4, Repeats: 1},
//		//{Name: "100-8-1", Nums: 100, Workers: 8, Repeats: 1},
//		//{Name: "100-16-1", Nums: 100, Workers: 16, Repeats: 1},
//		{Name: "100-32-1", Nums: 100, Workers: 32, Repeats: 1},
//		//{Name: "1000-1-1", Nums: 1000, Workers: 1, Repeats: 1},
//		//{Name: "1000-2-1", Nums: 1000, Workers: 2, Repeats: 1},
//		//{Name: "1000-4-1", Nums: 1000, Workers: 4, Repeats: 1},
//		//{Name: "1000-8-1", Nums: 1000, Workers: 8, Repeats: 1},
//		//{Name: "1000-16-1", Nums: 1000, Workers: 16, Repeats: 1},
//		//{Name: "1000-32-1", Nums: 1000, Workers: 32, Repeats: 1},
//	}
//
//	for _, c := range cases {
//		var r []int
//		generated := generator[int](1000)
//		b.Run(c.Name, func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				r = GeneProcessor(generated, 2, 1)
//			}
//			res = r
//		})
//	}
//}
