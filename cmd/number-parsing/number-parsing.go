package main

import (
	"fmt"
	"strconv"
)

// Parsing numbers from strings is a basic but common task in many programs; here's how to do it in Go...
// The build-in package strconv provides the number parsing
func main() {

	// With ParseFloat, this '64' tells us how many bits of precision to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println("float:", f)

	// For parse int, the '0' means infer the base from the string. 64 requires that the result fit in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println("int:", i)

	// ParseInt will recognize hex-formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println("int:", d)

	// A ParseUInt is also available
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println("uns:", u)

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println("atoi:", k)

	// Parse functions return an erro on bad input
	_, e := strconv.Atoi("what")
	fmt.Println("err:", e)
}
