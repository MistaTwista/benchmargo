package genestat

import (
	"github.com/MistaTwista/generigo/internal/util"
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

func Processor(readersCount, writersCount int) (uint64, uint64) {
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

	rand5 := util.Rnd(5)
	for r := 0; r < readersCount; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand5(),
					resp: make(chan int),
				}

				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	rand100 := util.Rnd(100)
	for w := 0; w < writersCount; w++ {
		go func() {
			for {
				write := writeOp{
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
