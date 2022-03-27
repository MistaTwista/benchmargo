package genevari

import (
	"sync"
	"testing"
)

func TestJustProc(t *testing.T) {
	x := Pipe(Identity(2), Multiply, MultiplyWith(10), Plus)
	if x != 80 {
		t.Fail()
	}
}

var res int

func TestGenerator(t *testing.T) {
	mulBy2 := Multiplicator(2)
	resCh := Summer(Duplicator(mulBy2(Generator(10))))
	var wg sync.WaitGroup
	wg.Add(1)
	var result int
	go func() {
		defer wg.Done()
		for i := range resCh {
			result = i
		}
	}()

	wg.Wait()
	if result != 180 {
		t.Fail()
	}
}

func BenchmarkChanPipe(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
	}

	for _, c := range cases {
		var r int
		b.Run("just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mulBy2 := Multiplicator(2)
				resCh := Summer(Duplicator(mulBy2(Generator(c.Nums))))

				r = <-resCh
			}

			res = r
		})
	}
}

func BenchmarkJustProc(b *testing.B) {
	cases := []struct {
		Name string
		Nums int
	}{
		{Name: "10", Nums: 10},
	}

	for _, c := range cases {
		var r int

		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := Pipe(Identity(c.Nums), Multiply, MultiplyWith(c.Nums), Plus)
				r = x
			}

			res = r
		})
	}
}
