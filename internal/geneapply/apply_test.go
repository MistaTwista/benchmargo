package geneapply

import (
	"testing"
)

func TestApplyJ(t *testing.T) {
	x := Apply(Identity(2), Multiply, MultiplyWith(10), Twice)
	if x != 80 {
		t.Fail()
	}
}

var res int

func BenchmarkApplyJ(b *testing.B) {
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
				x := Apply(Identity(c.Nums), Multiply, MultiplyWith(c.Nums), Twice)
				r = x
			}

			res = r
		})
	}
}
