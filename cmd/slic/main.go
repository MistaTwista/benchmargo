package main

import (
	"log"

	"github.com/MistaTwista/generigo/internal/geneslic"
)

func main() {
	res := geneslic.Uniq([]string{"one", "one", "two", "three"})
	log.Println(res)
	res = geneslic.GeneUniq([]string{"one", "one", "two", "three"})
	log.Println(res)
	resint := geneslic.GeneUniq([]int{1, 1, 2, 3})
	log.Println(resint)
}
