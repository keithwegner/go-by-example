package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A line filter is a common type of program that reads inputs on stdin, processes it, and then prints some derived
// result to stdout. grep and sed are common line filters.

// Here's an example line filter in Go that writes a capitalized version of all input text. You can use this pattern
// to write your own Go line filters.
func main() {
	// Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances
	// the scanner to the next token; which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	// scanner.Text returns the current token, here being the next line, from the input.
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		// Write out the line in uppercase.
		fmt.Println(ucl)
	}

	// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error", err)
		os.Exit(1)
	}

	// To try out the line filter, make a file with a few lowercase lines. Then use the line filter to get the
	// upper case lines.
	// >echo -e "hello\nworld" | go run line-filters.go
}
