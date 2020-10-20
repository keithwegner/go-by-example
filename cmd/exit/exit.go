package main

import (
	"fmt"
	"os"
)

// Program uses os.Exit to immediately exit with a given status

func main() {
	// defer will *not* be called when using os.Exit, so this Println will never be called
	defer fmt.Println("!")

	// Note that unlike, e.g., C, Go does not use an integer return value from main to indicate exit status.
	// If you'd like to exit with a non-zero status you should use os.Exit
	os.Exit(3)

	// If you run exit.go using 'go run', the exit will be picked up by go and printed.

	// By building an executing a binary you can see the status in the terminal.
}
