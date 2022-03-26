package genevari

import (
	"testing"
)

func TestProc(t *testing.T) {
	x := Pipe(Identity(2), Multiply, MultiplyWith(10), Plus)
	if x != 80 {
		t.Fail()
	}
}

var res int

func BenchmarkProc(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
	}

	for _, c := range cases {
		var r int
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := Pipe(Identity(c.Nums), Multiply, MultiplyWith(c.Nums), Plus)
				r = x
			}

			res = r
		})
	}
}

func BenchmarkGProc(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
	}

	for _, c := range cases {
		var r int
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := GPipe[int](GIdentity(c.Nums), GMultiply[int], GMultiplyWith(c.Nums), GPlus[int])
				r = x
			}

			res = r
		})
	}
}
