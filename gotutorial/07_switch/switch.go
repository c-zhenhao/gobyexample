// switch statements express conditionals across many branches

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// you cna use commas to seperate multiple expressions in the same case statement. We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ladies and gentlemen, its the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// a type switch compare types instead of values. you can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	fmt.Println("true")
	whatAmI(true)

	fmt.Println("1")
	whatAmI(1)

	fmt.Println("\"hey\"")
	whatAmI("hey")

	fmt.Println("1.0")
	whatAmI(1.0)

	fmt.Println("[]int{1, 2, 3}")
	whatAmI([]int{1, 2, 3})

	fmt.Println("map[string]int{\"a\": 1, \"b\": 2}")
	whatAmI(map[string]int{"a": 1, "b": 2})
}
