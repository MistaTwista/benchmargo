//go:build generics

package geneapply

import (
	"testing"
)

func TestApplyG(t *testing.T) {
	x := GApply(GIdentity(2), GMultiply[int], GMultiplyWith(10), GPlus[int])
	if x != 80 {
		t.Fail()
	}
}

func TestApplyGStrings(t *testing.T) {
	x := GApply(GIdentity("123"), GPlus[string])
	if x != "123123" {
		t.Fail()
	}
}

func BenchmarkApplyG(b *testing.B) {
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
				x := GApply[int](GIdentity(c.Nums), GMultiply[int], GMultiplyWith(c.Nums), GPlus[int])
				r = x
			}

			res = r
		})
	}
}
