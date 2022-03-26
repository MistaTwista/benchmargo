package genestat

import (
	"github.com/MistaTwista/generigo/internal/util"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func Processor() (uint64, uint64) {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	randomchik := util.Rnd(5)
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  randomchik(),
					resp: make(chan int),
				}

				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}

				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	return readOpsFinal, writeOpsFinal
}

type greadOp[T util.Numbers] struct {
	key  T
	resp chan T
}
type gwriteOp[T util.Numbers] struct {
	key  T
	val  T
	resp chan bool
}

func GProcessor[T util.Numbers]() (uint64, uint64) {
	var readOps uint64
	var writeOps uint64

	reads := make(chan greadOp[T])
	writes := make(chan gwriteOp[T])

	go func() {
		var state = make(map[T]T)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	rand5 := util.Grnd[T](5)
	rand100 := util.Grnd[T](100)
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := greadOp[T]{
					key:  rand5(),
					resp: make(chan T),
				}

				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := gwriteOp[T]{
					key:  rand5(),
					val:  rand100(),
					resp: make(chan bool),
				}

				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	return readOpsFinal, writeOpsFinal
}
