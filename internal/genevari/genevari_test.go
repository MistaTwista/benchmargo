//go:build generics

package genevari

import (
	"sync"
	"testing"
)

func TestGenetProc(t *testing.T) {
	x := GPipe(GIdentity(2), GMultiply[int], GMultiplyWith(10), GPlus[int])
	if x != 80 {
		t.Fail()
	}
}

func TestGeneGenerator(t *testing.T) {
	mulBy2 := GeneMultiplicator(2)
	resCh := GeneSummer(GeneDuplicator(mulBy2(GeneGenerator[int](10))))
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

func BenchmarkGeneProc(b *testing.B) {
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
	}
}

func BenchmarkGeneChanPipe(b *testing.B) {
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
	}
}
