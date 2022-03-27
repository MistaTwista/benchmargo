//go:build generics

package genechan

import "github.com/MistaTwista/generigo/internal/util"

func gchanputin[T util.Numbers](some T, ch chan T) {
	ch <- some
}

func GProc[T util.Numbers](n T) T {
	resCh := make(chan T)
	go gchanputin(n, resCh)

	return <-resCh
}

func ProcWithG(num int) int {
	resCh := make(chan int)

	go gchanputin(num, resCh)

	return <-resCh
}
