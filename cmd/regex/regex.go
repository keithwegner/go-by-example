package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go offers builtin support for regex. Here are some examples of common regexp-related tasks in Go.

func main() {
	// Tests whether a pattern matches a string.
	m, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(m)

	// Above we used a string pattern directly, but for other tasks you'll need to Compile an optimized Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. here's a match test like we saw earlier.
	fmt.Println(r.MatchString("peach"))

	// This finds the match for the regexp.
	fmt.Println(r.FindString("peach punch"))

	// This also finds the first match, but returns the start and end indexes for the match instead of the matching text.
	fmt.Println(r.FindStringIndex("peach punch"))

	// Submatch variants include info about both the whole-pattern matches and the submatches within those matches.
	// For example, this will return info for both p([a-z]+)ch and ([a-z]+).
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The All variants of these functions apply to all matches in the input, not just the first. For example to find all
	// matches for a regexp.
	fmt.Println(r.FindAllString("peach punch ping", -1))

	// These All variants are available for the other functions we saw above.
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Providing a non-negative integer as the second argument to these functions will limit the number of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Our examples had string ars and used names like MatchString. We can also provide []byte args and drop String
	// from the function name.
	fmt.Println(r.Match([]byte("peach")))

	// When creating global vars with regular expressions you can use the MustCompile variation of Compile. MustCompile
	// panics instead of returning an error, which makes it safer to use for global vars.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// The regexp package can also be used to replace subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The Func variant allows you to transform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
