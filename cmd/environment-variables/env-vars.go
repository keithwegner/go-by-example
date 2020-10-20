package main

import (
	"fmt"
	"os"
	"strings"
)

// Environment variables are a universal mechanism for conveying configuration
// information to Unix programs. Let's look at how to set, get, and list env vars.

func main() {
	// To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv.
	// This will return an empty string if the key isn't present in the env.
	f, b := "FOO", "BAR"

	os.Setenv(f, "1")
	fmt.Println(f, os.Getenv(f))
	fmt.Println(b, os.Getenv(b))

	fmt.Println()

	// Use os.Environ to list all key/value pairs in the env. This returns a slice of
	// strings in the form of KEY=value. You can strings.SplitN them to get the key
	// and value. Here we print all the keys.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", -1)
		fmt.Println(pair[0])
	}
	// Running the progrma shows we pick up the value for FOO that we sent in the
	// program, but that BAR is empty.

	// If we set BAR in the env first, the running program picks up that value. E.g.,
	// BAR=2 go run env-vars.go
	// 	FOO 1
	// 	BAR 2
}
