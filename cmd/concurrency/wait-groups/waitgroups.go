package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple Goroutines to finish, we can use a wait group

//  This is the function we'll run in every Goroutine. Note that a WaitGroup must be passed to
// functions by pointer
func worker(id int, wg *sync.WaitGroup) {

	// on return, notify the WaitGroup that we're done
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	//  The WaitGroup is used to wait for all the Goroutines launched here to finish
	var wg sync.WaitGroup

	// Launch several Goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they're done
	wg.Wait()
}
