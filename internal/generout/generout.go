package genechan

import "github.com/MistaTwista/generigo/internal/util"

func chanputin(ch chan int, some int) {
	ch <- some
}

func Proc(num int) int {
	resCh := make(chan int)

	go chanputin(resCh, num)

	return <-resCh
}

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

func chanputinInt(n interface{}, ch chan interface{}) {
	ch <- n
}

func ProcOnInt(some interface{}) int {
	resCh := make(chan interface{})

	go chanputinInt(some, resCh)
	x := <-resCh
	num, ok := x.(int)
	if !ok {
		return 0
	}

	return num
}
