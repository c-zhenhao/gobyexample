// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

package main

import "fmt"

func main() {
	var a = "initial" // var declares 1 or more variables
	fmt.Println(a)

	var b, c int = 1, 2 // can declare multiple variables at once
	fmt.Println(b, c)

	var d = true // go will infer the type of initialized variables
	fmt.Println(d)

	var e int // variables declared without a corresponding initalisation are zero-valued. for example, the zero value for an int is 0.
	fmt.Println(e)

	// the := syntax is shorthand for declaring and intializing a variable, e.g. for var f string = "apple" in this case.
	f := "apple" // := is only used in function bodies?
	// var f string = "apple" // this is equivalent to the above
	fmt.Println(f)

	var g string = "banana"
	fmt.Println(g)

	var h string
	fmt.Println(h)
}
