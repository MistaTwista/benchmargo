package genechan

import "testing"

func TestProc(t *testing.T) {
	r := Proc(42)
	if r != 42 {
		t.Fail()
	}
}

func TestGProc(t *testing.T) {
	r := GProc(42)
	if r != 42 {
		t.Fail()
	}
}

var res int

func BenchmarkProcessor(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = Proc(42)
	}
	res = r
}

func BenchmarkGProcessor(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = GProc(42)
	}
	res = r
}

func BenchmarkProcessorWithG(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = ProcWithG(42)
	}
	res = r
}

func BenchmarkProcessorOnInt(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = ProcOnInt(42)
	}
	res = r
}
