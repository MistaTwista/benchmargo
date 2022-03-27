//go:build generics

package genechan

import "testing"

func TestGeneProc(t *testing.T) {
	r := GProc(42)
	if r != 42 {
		t.Fail()
	}
}

func BenchmarkGeneProcessor(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = GProc(42)
	}
	res = r
}

func BenchmarkGeneProcessorWithG(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = ProcWithG(42)
	}
	res = r
}
