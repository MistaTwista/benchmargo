package genelast

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func TestLastJ(t *testing.T) {
	r := Last(42, 44)
	if r != 44 {
		t.Fail()
	}
}

var res int
var resAny interface{}

func BenchmarkLastJ(b *testing.B) {
	cases := []struct {
		Name  string
		Count int
	}{
		{Name: "10", Count: 10},
		{Name: "100", Count: 100},
	}

	for _, c := range cases {
		nums := util.GenerateRandoms(c.Count)
		numsAsInterfaces := util.IntsToInterface(nums...)

		b.Run(c.Name, func(b *testing.B) {
			var r int
			for i := 0; i < b.N; i++ {
				r = Last(nums...)
			}
			res = r
		})

		b.Run(c.Name+"-interface", func(b *testing.B) {
			var r interface{}
			for i := 0; i < b.N; i++ {
				r = LastInterface(numsAsInterfaces...)
			}
			resAny = r
		})
	}
}
