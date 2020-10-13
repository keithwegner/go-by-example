package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for mataining state in Go is communication over channels. We saw this example with
// Worker Pools. There are a few other options for maintaining state, though. Here, we will look at using the
// sync/atomic package for Atomic Counters accessed by multiple Goroutines

func main() {

	// We'll use an unsigned integer to represent our (always positive) counter
	var ops uint64

	// A WaitGroup will help us wait for all Goroutines to finish their work
	var wg sync.WaitGroup

	// We'll start 50 Goroutines that each increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// To atomically increment the counter, we use AddUint64, giving it the memory
		// address of our ops counter.
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Wait until all of the Goroutines are done
	wg.Wait()

	// It's safe to access ops now because we know no other Goroutine is writing to it. Reading atomics safely
	// while they are being updated is also positive using functions like atomic.LoadUint64
	fmt.Println("ops:", ops)

	// We expect to get exactly 50,000 operations. Had we used the non-atomic ops++ to increment the counter,
	// we'd likely get a different number, changing between runs, because the Goroutines would interfere with
	// one another. Moreover we'd get data race failures when running with the -race flag.

}
