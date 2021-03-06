package geneplus

import (
	"fmt"
	"testing"

	"github.com/MistaTwista/benchmargo/internal/util"
)

func TestJustAddTwoIntPlus(t *testing.T) {
	cases := []struct {
		Name   string
		A      int
		B      int
		Result int
	}{
		{Name: "ints", A: 2, B: 2, Result: 4},
	}

	for _, c := range cases {
		t.Run(c.Name+" interface", func(t *testing.T) {
			res, _ := addWithInterface(c.A, c.B)
			result, _ := res.(int)
			if c.Result != result {
				t.Fatal("!=")
			}
		})
		t.Run(c.Name+" reflect", func(t *testing.T) {
			res, _ := addReflected(c.A, c.B)
			result, _ := res.(int)
			if c.Result != result {
				t.Fatal("!=")
			}
		})
		t.Run(c.Name+" just", func(t *testing.T) {
			res := addInt(c.A, c.B)
			if c.Result != res {
				t.Fatal("!=")
			}
		})
	}
}

func TestJustAddIntVariadicPlus(t *testing.T) {
	cases := []struct {
		Name   string
		List   []interface{}
		Result interface{}
	}{
		{Name: "ints", List: []interface{}{1, 2, 3, 4, 5}, Result: 15},
		{Name: "strings", List: []interface{}{"hello", "world", "whats", "up"}, Result: "helloworldwhatsup"},
		{Name: "floats", List: []interface{}{1.1, 1.2, 1.3, 1.4}, Result: 5.0},
	}

	for _, c := range cases {
		t.Run(c.Name+" interface", func(t *testing.T) {
			res, err := addWithInterfaceVari(c.List...)
			if err != nil {
				t.Fatal(err)
			}

			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})

		t.Run(c.Name+" reflect", func(t *testing.T) {
			res, err := addReflectedVari(c.List...)
			if err != nil {
				t.Fatal(err)
			}

			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})
	}
}

func TestJustAddIntsVariadicPlus(t *testing.T) {
	cases := []struct {
		Name   string
		List   []int
		Result int
	}{
		{Name: "ints", List: []int{1, 2, 3, 4, 5}, Result: 15},
	}

	for _, c := range cases {
		t.Run(c.Name+" int", func(t *testing.T) {
			res := addIntVari(c.List...)
			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})
	}
}

func TestJustAddFloats64VariadicPlus(t *testing.T) {
	cases := []struct {
		Name   string
		List   []float64
		Result float64
	}{
		{Name: "floats", List: []float64{1.1, 1.2, 1.3, 1.4}, Result: 5.0},
	}

	for _, c := range cases {
		t.Run(c.Name+" just", func(t *testing.T) {
			res := addFloat64Vari(c.List...)
			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})
	}
}

func TestJustAddStringVariadicPlus(t *testing.T) {
	cases := []struct {
		Name   string
		List   []string
		Result string
	}{
		{Name: "strings", List: []string{"hello", "world", "whats", "up"}, Result: "helloworldwhatsup"},
	}

	for _, c := range cases {
		t.Run(c.Name+" just", func(t *testing.T) {
			res := addStringVari(c.List...)
			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})

		t.Run(c.Name+" just buffered", func(t *testing.T) {
			res := addStringBufVari(c.List...)
			if c.Result != res {
				t.Fatal(fmt.Sprintf("want: %v, got: %v", c.Result, res))
			}
		})
	}
}

var resultInt int

func BenchmarkAddTwoIntJ(b *testing.B) {
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
		b.Run(c.Name+"-interface", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addWithInterface(c.A, c.B)
				r = res.(int)
			}
			resultInt = r
		})

		b.Run(c.Name+"-reflect", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addReflected(c.A, c.B)
				r = res.(int)
			}
			resultInt = r
		})

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addInt(c.A, c.B)
				r = res
			}
			resultInt = r
		})
	}
}

var resultFloat float64

func BenchmarkAddTwoFloat64J(b *testing.B) {
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
		b.Run(c.Name+"-interface", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addWithInterface(c.A, c.B)
				r = res.(float64)
			}
			resultFloat = r
		})

		b.Run(c.Name+"-reflect", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addReflected(c.A, c.B)
				r = res.(float64)
			}
			resultFloat = r
		})

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addFloat64(c.A, c.B)
				r = res
			}
			resultFloat = r
		})
	}
}

var resultString string

func BenchmarkAddTwoStringJ(b *testing.B) {
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
		b.Run(c.Name+"-interface", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addWithInterface(c.A, c.B)
				r = res.(string)
			}
			resultString = r
		})

		b.Run(c.Name+"-reflect", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addReflected(c.A, c.B)
				r = res.(string)
			}
			resultString = r
		})

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addString(c.A, c.B)
				r = res
			}
			resultString = r
		})

		b.Run(c.Name+"-buf", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addStringBuf(c.A, c.B)
				r = res
			}
			resultString = r
		})
	}
}

func BenchmarkAddIntsJ(b *testing.B) {
	cases := []struct {
		Name string
		List []int
	}{
		{Name: "ints", List: util.RepeatInts(10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)},
	}

	for _, c := range cases {
		var r int
		interfacedList := util.IntsToInterface(c.List...)

		b.Run(c.Name+"-interface", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addWithInterfaceVari(interfacedList...)
				r = res.(int)
			}

			resultInt = r
		})

		b.Run(c.Name+"-reflected", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addReflectedVari(interfacedList...)
				r = res.(int)
			}

			resultInt = r
		})

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addIntVari(c.List...)
				r = res
			}
			resultInt = r
		})
	}
}

func BenchmarkAddFloatsJ(b *testing.B) {
	cases := []struct {
		Name string
		List []float64
	}{
		{Name: "floats", List: util.RepeatFloats(10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)},
	}

	for _, c := range cases {
		var r float64
		interfacedList := util.FloatsToInterface(c.List...)

		b.Run(c.Name+"-interface", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addWithInterfaceVari(interfacedList...)
				r = res.(float64)
			}

			resultFloat = r
		})

		b.Run(c.Name+"-reflected", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addReflectedVari(interfacedList...)
				r = res.(float64)
			}

			resultFloat = r
		})

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addFloat64Vari(c.List...)
				r = res
			}
			resultFloat = r
		})
	}
}

func BenchmarkAddStringsJ(b *testing.B) {
	cases := []struct {
		Name string
		List []string
	}{
		{Name: "strings", List: util.RepeatStrings(10, "one", "two", "three", "four", "five")},
	}

	for _, c := range cases {
		var r string
		interfacedList := util.StringsToInterface(c.List...)

		b.Run(c.Name+"-interface", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addWithInterfaceVari(interfacedList...)
				r = res.(string)
			}

			resultString = r
		})

		b.Run(c.Name+"-reflected", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := addReflectedVari(interfacedList...)
				r = res.(string)
			}

			resultString = r
		})

		b.Run(c.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addStringVari(c.List...)
				r = res
			}
			resultString = r
		})

		b.Run(c.Name+"-buffered", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res := addStringBufVari(c.List...)
				r = res
			}
			resultString = r
		})
	}
}
