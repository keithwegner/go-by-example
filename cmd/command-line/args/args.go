package main

import (
	"fmt"
	"os"
)

// Command line arguments are a common way to parameterize execution of programs. For
// example, 'go run hello.go' uses 'run' and 'hello.go' arguments to the go program.

func main() {
	// os.Args provides access to raw command line arguments. Note that the first value
	// in this slice is the path to the program, and os.Args[1:] holds the args to the
	// program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// To experiment with command line arguments it's best to build a binary
	// with 'go build' first.
	//   go build command-line-args.go
	//   ./command-line-args.go a b c d
}
