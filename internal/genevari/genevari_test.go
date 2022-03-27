package genevari

import (
	"sync"
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
		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := GPipe[int](GIdentity(c.Nums), GMultiply[int], GMultiplyWith(c.Nums), GPlus[int])
				r = x
			}

			res = r
		})

		b.Run(c.Name+" just", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := Pipe(Identity(c.Nums), Multiply, MultiplyWith(c.Nums), Plus)
				r = x
			}

			res = r
		})
	}
}

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
		b.Run("generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mulBy2 := GeneMultiplicator(2)
				resCh := GeneSummer(GeneDuplicator(mulBy2(GeneGenerator[int](c.Nums))))

				r = <-resCh
			}

			res = r
		})

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
