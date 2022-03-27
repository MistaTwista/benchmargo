package genepipe

import (
	"testing"
)

var res int

func TestPipeJ(t *testing.T) {
	mulBy2 := Multiplicator(2)
	result := <-Add(Flat(Duplicator(mulBy2(Generator(2, []int{0, 1, 2, 3}...)))))

	if result != 48 {
		t.Fail()
	}
}

func BenchmarkPipeJ(b *testing.B) {
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
				mulBy2 := Multiplicator(2)
				r = <-Add(Flat(Duplicator(mulBy2(Generator(c.Nums, c.List...)))))
			}

			res = r
		})
	}
}
