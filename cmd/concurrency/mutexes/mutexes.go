package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// In our Atomic Counters example we saw how to manage simple counter state using atomic operations.
// For more complex state we can use a mutex to safely access data across multiple Goroutines

// (Useful article on the runtime errors reported when a map is modifified unsafely by concurrent execution:
// https://medium.com/a-journey-with-go/go-concurrency-access-with-maps-part-iii-8c0a0e4eb27e)

func main() {
	// For our example the state will be a map
	var state = make(map[int]int)

	// This mutex will synchronize access to state
	var mutex = &sync.Mutex{}

	// We'll keep track of how many read and write operations we perform
	var readOps uint64
	var writeOps uint64

	// here we start 100 Goroutines to executes repeated reads aagainst the state, once per millisecond in each
	// Goroutine
	for r := 0; r < 2; r++ {
		go func() {
			total := 0

			// For each read we pick a key to access, Lock() the mutex to ensure exclusive access to the state,
			// Unlock() the mutex, and increment the readOps count
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// Wait a bit between reads
				time.Sleep(time.Millisecond)

			}
		}()
	}

	// We'll also start 10 Goroutines to simulate writes, using the same pattern we did for the reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// Pick a key and value, Lock() the mutex to again ensure exlusive acces to the state,
				// update the key with the value, Unlock() the mutex, and increment the writeOps count
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)

				// Wait a bit between reads
				time.Sleep(time.Millisecond)

			}
		}()
	}

	//  Let the 10 Goroutines wok on the state and mutex for a second
	time.Sleep(time.Second)

	// Take and report final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// With a final lock of state, show how it ended up
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

	// Running the program shows that we execute about 90,000 total operations against our
	// mutex-synchronized state
}
