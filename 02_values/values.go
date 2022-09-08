// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

package main

import "fmt"

func main() {
	// Strings, which can be added together with +.
	fmt.Println("Go" + "lang")		// prints "go lang"
	
	// integers and floats
	fmt.Println("1+1 =", 1+1)		// prints 2
	fmt.Println("7.0/3.0 =", 7.0/3.0)	// expect a float, so 2.3333333
	s := fmt.Sprintf("%.2f", 7.0/3.0)	// Sprintf returns as string
	fmt.Printf(s)

	fmt.Println("--booleans--")
	fmt.Println(true && false)	// false
	fmt.Println(true || false)	// true
	fmt.Println(!true)	// false
}