//go:build generics

package genetochan

import (
	"github.com/MistaTwista/benchmargo/internal/util"
	"testing"
)

func TestToChan(t *testing.T) {
	cases := []struct {
		Name  string
		Items []int
	}{
		{Name: "10", Items: []int{1, 2, 3}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := make([]<-chan int, 0, len(c.Items))
			var r int
			for _, n := range c.Items {
				n := n
				res = append(res, GToChan[int](func() int {
					return n * 2
				}))
			}

			for x := range GMix[int](res...) {
				r += x
			}
			if r != 12 {
				t.Fatalf("res %d", r)
			}
		})
	}
}

func TestShit(t *testing.T) {
	one, two, three, four := GToChan[int](With(5, Ticker)),
		GToChan[int](With(5, Ticker)),
		GToChan[int](With(5, Ticker)),
		GToChan[string](With("kernel", Pucker))

	res := make([]<-chan int, 0)
	res = append(res, GToChan[int](With(5, Ticker)))
	res = append(res, GToChan[int](With(5, Ticker)))
	res = append(res, GToChan[int](With(5, Ticker)))
	result := 0
	for num := range GMix(res...) {
		result += num
	}
	t.Log(result)
	t.Log(<-one, <-two, <-three, <-four)
}

func BenchmarkToChanG(b *testing.B) {
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
		b.Run(c.Name+" typed arg", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := make([]<-chan int, 0, len(c.Items))
				for _, n := range c.Items {
					n := n
					res = append(res, GToChan[int](func() int {
						return n
					}))
				}

				for x := range GMix[int](res...) {
					r = x
				}
			}
		})

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := make([]<-chan int, 0, len(c.Items))
				for _, n := range c.Items {
					n := n
					res = append(res, GToChanAny[int](func() int {
						return n
					}))
				}

				for x := range GMix[int](res...) {
					r = x
				}
			}
		})
		result = r
	}
}
