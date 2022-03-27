package genetochan

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
	"time"
)

func Ticker(n int) int {
	time.Sleep(5 * time.Second)
	return n
}

func Pucker(shit string) string {
	time.Sleep(6 * time.Second)
	return shit + "42"
}

var result int

func BenchmarkToChanJ(b *testing.B) {
	cases := []struct {
		Name  string
		Items []int
	}{
		{Name: "3", Items: []int{1, 2, 3}},
		{Name: "10", Items: util.GenerateRandoms(10)},
		{Name: "100", Items: util.GenerateRandoms(100)},
		{Name: "1000", Items: util.GenerateRandoms(1000)},
		{Name: "10000", Items: util.GenerateRandoms(10000)},
	}

	for _, c := range cases {
		var r int
		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := make([]<-chan int, 0, len(c.Items))
				for _, n := range c.Items {
					n := n
					res = append(res, ToChan(func() int {
						return n
					}))
				}

				for x := range Mix(res...) {
					r = x
				}
			}
		})
		result = r
	}
}
