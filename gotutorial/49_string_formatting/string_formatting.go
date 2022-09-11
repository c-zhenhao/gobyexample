// Go offers excellent support for string formatting in the printf tradition. here are some examples of common string formatting tasks
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go offers several printing "verbs" designed to format general Go values. For example, this prints an instance of our point struct
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// if the value is a struct, the %+v variant will include the struct's field names
	fmt.Printf("struct2: %+v\n", p)

	// the %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value
	fmt.Printf("struct3: %#v\n", p)

	// to print the type of a value, use %T
	fmt.Printf("type: %T\n", p)

	// formatting booleans is straight-forward
	fmt.Printf("bool: %t\n", true)

	// there are many options for formatting integers. use %d for standard, base-10 formatting (d for dec)
	fmt.Printf("int: %d\n", 123)

	// this prints a binary representation
	fmt.Printf("bin: %b\n", 14)

	// this prints the character corresponding to the given integer
	fmt.Printf("char: %c\n", 33)

	// %x provides hex encoding
	fmt.Printf("hex: %x\n", 456)

	// there are also several formatting options for floats. %f prints basic decimal formatting
	fmt.Printf("float1: %f\n", 78.9)

	// %e and %E format the float in (slightly different versions of) scientific notation
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// for basic string printing use %s
	fmt.Printf("string1: %s\n", "\"string\"")

	// for double-quoted strings use %q
	fmt.Printf("string2: %q\n", "\"string\"")

	// as with integers, %x renders the string in base-16, with two output characters per byte of input
	fmt.Printf("string3: %x\n", "hex this")

	// to print a representation of a pointer, use %p
	fmt.Printf("pointer: %p\n", &p)

	// when formatting numbers you will often want to control the width and precision of the resulting figure. to specify the width of an integer, use a number after the % in the verb. by default the result will be right-justified and padded with spaces
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// you can also specify the width of printed floats, though usually you'll also want to restrict the decimal precision at the same time with the width.precision syntax
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// to left-justify, use the - flag
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// you may also want to control width when formatting strings, especially to ensure that they align in table-like output. for basic right-justified width
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// to left-justify use the - flag as with numbers
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// so far we've seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere
	s := fmt.Sprintf("string: %s", "string")
	fmt.Println(s)

	// you can format+print to io.Writers other than os.Stdout using Fprintf
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
