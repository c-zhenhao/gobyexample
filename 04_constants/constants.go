//  Go supports constants of character, string, boolean, and numeric values.

package main

import (
	"fmt"
	"math"
)

// constant expressions perform arithmetic with a arbitrary precision
const s string = "constant"

func main() {
	fmt.Println(s)
	
	// a const statment can appear anywhere a var statement can
	const n = 50000000

	// numeric constant has no type until its given one, such as one by an explicit conversion
	const d = 3e20 / n 
	fmt.Println(d)

	// a number can be given type by using it in a context that requires one, such as a variable assignment or function call
	fmt.Println(int64(d))

	// for example, here math.Sin expects float64
	fmt.Println(math.Sin(n))
}