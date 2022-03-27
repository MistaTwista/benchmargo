//go:build generics

package geneplus

import (
	"fmt"
	"testing"

	"github.com/MistaTwista/generigo/internal/util"
)

func TestGeneAddInt(t *testing.T) {
	cases := []struct {
		Name   string
		A      int
		B      int
		Result int
	}{
		{Name: "ints", A: 2, B: 2, Result: 4},
	}

	for _, c := range cases {
		t.Run(c.Name+"generics", func(t *testing.T) {
			res := addGenerics(c.A, c.B)
			if c.Result != res {
				t.Fatal("!=")
			}
		})
	}
}

func TestGeneAddIntsVariadic(t *testing.T) {
	cases := []struct {
		Name   string
		List   []int
		Result int
	}{
		{Name: "ints", List: []int{1, 2, 3, 4, 5}, Result: 15},
	}

	for _, c := range cases {
		t.Run(c.Name+" generic", func(t *testing.T) {
			res := addGenericsVari(c.List...)
			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})
	}
}

func TestGeneAddFloats64Variadic(t *testing.T) {
	cases := []struct {
		Name   string
		List   []float64
		Result float64
	}{
		{Name: "floats", List: []float64{1.1, 1.2, 1.3, 1.4}, Result: 5.0},
	}

	for _, c := range cases {
		t.Run(c.Name+" generic", func(t *testing.T) {
			res := addGenericsVari(c.List...)
			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})
	}
}

func TestGeneAddStringVariadic(t *testing.T) {
	cases := []struct {
		Name   string
		List   []string
		Result string
	}{
		{Name: "strings", List: []string{"hello", "world", "whats", "up"}, Result: "helloworldwhatsup"},
	}

	for _, c := range cases {
		t.Run(c.Name+" generic", func(t *testing.T) {
			res := addGenericsVari(c.List...)
			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})
	}
}

func BenchmarkGeneAddInt(b *testing.B) {
	cases := []struct {
		Name   string
		A      int
		B      int
		Result int
	}{
		{Name: "ints", A: 2, B: 2, Result: 4},
	}

	for _, c := range cases {
		var r int
		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addGenerics(c.A, c.B)
				r = res
			}
			resultInt = r
		})
	}
}

func BenchmarkGeneAddFloat64(b *testing.B) {
	cases := []struct {
		Name   string
		A      float64
		B      float64
		Result float64
	}{
		{Name: "floats", A: 2, B: 2, Result: 4},
	}

	for _, c := range cases {
		var r float64
		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addGenerics(c.A, c.B)
				r = res
			}
			resultFloat = r
		})
	}
}

func BenchmarkGeneAddString(b *testing.B) {
	cases := []struct {
		Name   string
		A      string
		B      string
		Result string
	}{
		{Name: "strings", A: "2", B: "2", Result: "22"},
	}

	for _, c := range cases {
		var r string
		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addGenerics(c.A, c.B)
				r = res
			}
			resultString = r
		})
	}
}

func BenchmarkGeneAddIntVari(b *testing.B) {
	cases := []struct {
		Name string
		List []int
	}{
		{Name: "ints", List: util.Repeat(10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)},
	}

	for _, c := range cases {
		var r int
		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addGenericsVari(c.List...)
				r = res
			}
			resultInt = r
		})
	}
}

func BenchmarkGeneAddFloatVari(b *testing.B) {
	cases := []struct {
		Name string
		List []float64
	}{
		{Name: "floats", List: util.Repeat[float64](10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)},
	}

	for _, c := range cases {
		var r float64
		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addGenericsVari(c.List...)
				r = res
			}
			resultFloat = r
		})
	}
}

func BenchmarkGeneAddStringsVari(b *testing.B) {
	cases := []struct {
		Name string
		List []string
	}{
		{Name: "strings", List: util.Repeat[string](10, "one", "two", "three", "four", "five")},
	}

	for _, c := range cases {
		var r string
		b.Run(c.Name+" generics", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addGenericsVari(c.List...)
				r = res
			}
			resultString = r
		})
	}
}
