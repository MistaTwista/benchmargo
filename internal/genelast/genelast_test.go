//go:build generics

package genelast

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func TestLastG(t *testing.T) {
	r := GLast(42, 43)
	if r != 43 {
		t.Fail()
	}
}

func TestLastInterface(t *testing.T) {
	r := GLast[interface{}]([]interface{}{1, "one", []int{4, 3, 2}}...)
	resultInterface = r
	t.Logf("%v", resultInterface)
}

var resultInterface interface{}

func BenchmarkLastG(b *testing.B) {
	cases := []struct {
		Name  string
		Count int
	}{
		{Name: "10", Count: 10},
		{Name: "100", Count: 100},
	}

	for _, c := range cases {
		nums := util.GenerateRandoms(c.Count)

		b.Run(c.Name+"-just-with-generic", func(b *testing.B) {
			var r int
			for i := 0; i < b.N; i++ {
				r = LastWithG(nums...)
			}
			res = r
		})

		b.Run(c.Name, func(b *testing.B) {
			var r int
			for i := 0; i < b.N; i++ {
				r = GLast(nums...)
			}
			res = r
		})

		numsAsInterface := util.IntsToInterface(nums...)
		b.Run(c.Name+" interface{}", func(b *testing.B) {
			var r interface{}
			for i := 0; i < b.N; i++ {
				r = GLast[interface{}](numsAsInterface...)
			}
			resultInterface = r
		})
	}
}
