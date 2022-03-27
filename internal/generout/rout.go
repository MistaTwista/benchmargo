package genechan

func chanputin(ch chan int, some int) {
	ch <- some
}

func Proc(num int) int {
	resCh := make(chan int)

	go chanputin(resCh, num)

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
