package main

import (
	"fmt"
	"sort"
)

// Sometimes we'll want to sort a collection by something other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of alphabetically. Here is an example of custom sorts in Go.

// In order to sort by a custom function in Go, we need a corresponding type. Here we've created a byLength type
// that is just an  alias for the builtin []string type.
type byLength []string

// Implement the sort.Interface - Len, Less and Swap - on our type so we can use the Sort package's generic
// Sort function. Len and Swap will usually be similar across types, and Less will hold the actual custom
// sorting logic. In our case, we want to sort in order of increasing string length, so we will use len(s[i])
//  and len(s[j]) here.

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	bl := byLength(fruits)
	sort.Sort(bl)
	fmt.Println(fruits)
}
