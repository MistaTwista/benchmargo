//go:build generics

package genepipe

import (
	"testing"
)

func TestPipeG(t *testing.T) {
	mulBy2 := GeneMultiplicator(2)
	result := <-GeneSummer(GeneFlat(GeneDuplicator(mulBy2(GeneGenerator(2, []int{0, 1, 2, 3}...)))))

	if result != 48 {
		t.Fail()
	}
}

func BenchmarkPipeG(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
		List []int
	}{
		{Name: "10", Nums: 10, List: []int{0, 1, 2, 3}},
	}

	for _, c := range cases {
		var r int
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mulBy2 := GeneMultiplicator(2)
				r = <-GeneSummer(GeneFlat(GeneDuplicator(mulBy2(GeneGenerator(c.Nums, c.List...)))))
			}

			res = r
		})
	}
}
