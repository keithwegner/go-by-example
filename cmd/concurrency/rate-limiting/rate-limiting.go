package main

// Rate Limiting is an important mechanism for controlling resource utilization and maintaining
// quality of service. Go elegantlt supports rate limiting with Goroutines, channels and tickers

import (
	"fmt"
	"time"
)

func main() {
	// First we'll look at basic rate limiting.

	// Suppose we want to limit our handling of incoming requests. We'll serve these requests off
	// a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
		fmt.Println("Wrote request", i, "to requests channel")
	}
	close(requests)

	//  The limiter channel will receive a value every 200 milliseconds.
	//  This is the regulator in our rate limiting scheme.
	limiter := time.Tick(500 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request, we limit ourselves
	// to 1 request every 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Printf("\nBurst limiting\n\n")
	// We may want to allow short bursts of requests in our rate limiting scheme while preserving the
	// overall rate limit. We can accomplish this by buffering our limiter channel. This burstyLimiter
	// channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we'll try to add a new value to burstyLimiter, up to it's limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requets. The first 3 of these will benefit from the burst capability
	// of burstyLimiter
	const capacity int = 5
	requests = make(chan int, capacity)
	for i := 1; i <= capacity; i++ {
		requests <- i
		// fmt.Println("Wrote request ", i, "to requests")
	}
	close(requests) // close it so the below range doesn't wait forever
	for req := range requests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	// Running our first program we see the first batch of requests handled once every ~200 milliseconds

	// For the second batch of requests, we serve the first 3 immediately because of the burstable
	// rate limitting. Then we server the 2 remaining batches with ~200ms delays each.
}
