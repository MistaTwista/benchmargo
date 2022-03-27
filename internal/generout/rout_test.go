package genechan

import "testing"

func TestJustProc(t *testing.T) {
	r := Proc(42)
	if r != 42 {
		t.Fail()
	}
}

var res int

func BenchmarkJustProcessor(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = Proc(42)
	}
	res = r
}

func BenchmarkJustProcessorOnInt(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = ProcOnInt(42)
	}
	res = r
}
