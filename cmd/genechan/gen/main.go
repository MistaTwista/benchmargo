package main

import "github.com/MistaTwista/generigo/internal/genechan"

func main() {
	_ = genechan.GeneProcessor[int]([]int{1, 2, 3}, 2, 2)
}
