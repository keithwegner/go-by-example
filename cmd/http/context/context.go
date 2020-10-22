package main

import (
	"fmt"
	"net/http"
	"time"
)

// In the HTTP server example, we looked at setting up a simple HTTP server. HTTP servers are useful for demonstrating
// the usage of context.Context for controlling cancelling. A Context carries deadlines, cancellation signals, and
// other request-scoped values across API boundaries and Goroutines.
func hello(w http.ResponseWriter, req *http.Request) {
	// A context is created for each request by the net/http machinery, and is available with the Context() method.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait a few seconds before sending a reply to the client. This could simulate some work the server is doing.
	// While working, keep an eye on the context's Done() channel for a signal that we should cancel the work and
	// return as soon as possible.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello\n")
	case <-ctx.Done():
		// The context's Err() method returns an error that explains why the Done() channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalErr := http.StatusInternalServerError
		http.Error(w, err.Error(), internalErr)
	}
}

// As before, register the handler on the "/hello" route and start serving.
func main() {
	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}

// Run the server in the background, simulate a client request to /hello, hitting CTRL+C shortly after
// starting to signal cancellation.
// > go run context.go
// > curl localhost:8090/hello
// ^C
