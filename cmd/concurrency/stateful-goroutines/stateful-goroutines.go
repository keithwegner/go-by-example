package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In the previous example we used explicit locking with Mutexes to synchronize access to shared state across
// multiple Goroutines. Another option is to use the built-in synchronization features of Goroutines and channels
// to achieve the same result. This channel-based approach aligns with Go's ideas of sharing memory by
// communicating and having each piece of data owned by exactly one Goroutine

// In this example our state will be owned by a single Goroutine. This will guarantee that the data is never
// corrupted with concurrent access. In order to read or write that state, other Goroutines will send message
// to the owning Goroutine and receive corresponding replies. These ReadOp and WriteOp structs encapsulate
// those requests and a way for the owning Goroutine to respond.
type ReadOp struct {
	key  int
	resp chan int
}

type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// Similar to the Mutexes example, count how many operations we perform.
	var readOps uint64
	var writeOps uint64

	// The reads and writes channels will be used by other Goroutines to issue
	// read and write requests, respectively.
	reads := make(chan ReadOp)
	writes := make(chan WriteOp)

	// Here is the Goroutine that owns the state, which is a map as in the Mutex example, but now private to the
	// stateful Goroutine. This Goroutine repeatedly selects on  the reads and writes channels, responding to
	// requests as they arrive. A response is executed by first performing the requested operation and then
	// sending a value on the response channel 'resp' to indicate success (and the desired value in the case
	// of reads)
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

	// This starts 100 Goroutines to issue reads to the state-owning Goroutine via the 'reads' channel. Each
	// read requires constructing a ReadOp struct, sending it over the 'reads' channel, and the receiving
	// result over the provided 'resp' channel.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := ReadOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 write Goroutines as well, using a similar approach.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := WriteOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the Goroutines work for a second
	time.Sleep(time.Second)

	// Capture and report the operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps", writeOpsFinal)

	// This Goroutine-based approach was a bit more involved than the Mutex one. It might be useful in certain
	// cases, though. For example, where you have other channels involved or when managing multiple such
	// mutexes would be error-prone. You should use whichever approach feels the most natural, especially with
	// respect to understanding the correctness of your program.
}
