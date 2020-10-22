package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

// Flags are a common way to specify options for command line programs. Go provides a flag
// package supporting basic command line flag parsing. We'll use this package to implement our
// example command line program
func main() {
	// Basic flag declarations are available for string, integer and boolean options. Here we declare
	// a string flag 'word' witha  default value 'foo' and a short descprtion. The flag.String
	// function returns a string pointer (not a string value). We'll see how to use this pointer below
	wordPtr := flag.String("word", "foo", "a string")

	// This declares 'numb' and 'fork' flags, using a similar approach to the 'word' flag.
	numbPtr := flag.Int("number", 42, "an integer")
	boolPtr := flag.Bool("boolean", false, "a boolean")

	// It's also possible to declare an option that uses an existing var declared elsewhere in the
	// program. Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Fun with spinners
	w := wow.New(os.Stdout, spin.Get(spin.Dots), " Parsing flags")
	w.Start()
	time.Sleep(2 * time.Second)
	w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Got 'em!")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing
	flag.Parse()

	// Here, we'll just dump out the parsed options and any trailing positional arguments. Note that we
	// need to derefernce the pointers with, e.g., *wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
