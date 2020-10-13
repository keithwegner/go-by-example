package main

import (
	"fmt"
	"sort"
)

// Go's sort package implements sorting for builtins and user-defined types.
// We'll look at sorting for builtins first

func main() {
	// Sort methods are specific to the builtin type. Here is an example for strings.
	// Note that sorting is in-place, so it changes the given slice and doesn't return a new one
	strings := []string{"C", "A", "B"}
	sort.Strings(strings)
	fmt.Println("Strings:", strings)

	// An example sorting ints
	ints := []int{3, 1, 1000, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	//  We can also use sort to check if a slice is already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
