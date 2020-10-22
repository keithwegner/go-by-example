package main

import (
	"fmt"
	"net/http"
)

// Writing a basic HTTP server is easy using the net/http package.

// A fundamental concept in net/http servers is Handlers. A handler is an object implementing the http.Handler interface.
// A common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature.
func hello(w http.ResponseWriter, r *http.Request) {
	// Functions serving as handlers take an http.ResponseWriter and an http.Request as arguments. The response writer
	// is used to fill in the HTTP response. Here, our simple response is just "Hello\n".
	fmt.Fprintf(w, "Hello\n")
}

// This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them
// into the body.
func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Register the handlers on server routes using the http.HandleFunc convenience function.
	// It sets up the Default Router in the net/http package and takes a function as an argument.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finally, call ListenAndServe with a port and a handler. nil tells it to use the default router we just set up.
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}

	// Run the server in the background and access the routes
	// > go run server.go
	// > curl localhost:8090/hello
	// > curl localhost:8090/headers
}
