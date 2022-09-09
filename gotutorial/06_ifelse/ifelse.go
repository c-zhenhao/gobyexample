// branching with if and else in Go is straight-forward

package main

import "fmt"

func main() {

	// basic example
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else if 7 % 2 != 0 {
		fmt.Println("7 is odd")
	}

	// can have if statement without an else
	if 8 % 4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// a statement can precede conditionals; any variables declared in this are available in all branches
	if num := 0; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}

	// note that you dont need parentheses around conditions in Go, but braces are required

	// there is no tenerary in Go, so you will need to use a full if statement even for basic conditions
}