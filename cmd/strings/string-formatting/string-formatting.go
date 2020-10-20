package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {
	p := point{1, 2}

	//  Print an instance of the point struct
	pf("%v\n", p)

	// If the value is a struct, the %+v variant will include the struct field's names
	pf("%+v\n", p)

	// The %#v variant prints a Go syntax representation of the value; i.e., the source code snippet that
	// would produce the value
	pf("%#v\n", p)

	// Print the Type with %T
	pf("%T\n", p)

	// Formattinb booleans is straightforward
	pf("%t\n", true)

	// Many options for fomratting integers. Use %d for standard base-10 formatting
	pf("%d\n", 123)

	//  Prints a binary representation
	pf("%b\n", 14)

	// Prints the character corresponding to the given integer
	pf("%c\n", 33)

	// %x provides hex encoding
	pf("%x\n", 456)

	// Several formatting options for floats. For basic decimal formatting, use %f
	pf("%f\n", 78.9)

	// %e and %E format the float in slightly different versions of scientific notation
	pf("%e\n", 123400000.0)
	pf("%E\n", 123400000.0)

	// For basic string printing, use %s
	pf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q
	pf("%q\n", "\"string\"")

	// As with integers seen above, %x renders a string in base-16,
	// with two output characters per byte of input
	pf("%x\n", "hex this")

	// To print a representation of a pointer, use %p
	pf("%p\n", &p)

	// When formatting numbers, you'll often want to control the width and precision of the resulting figure.
	// To specifdy the widgth of an integer, use a number after the % in the verb. By default, the result
	// will be right-justified and padded with spaces.
	pf("|%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats, though usually you'll also wantt o restrict the
	// decimal precision at the same time with the width.precision syntax
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	pf("|%6s|%6s|\n", "foo", "b")

	// To left-justify use the - flag as with numbers.
	pf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without printing it anywhere.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// You can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
