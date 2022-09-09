// in Go, its idiomatic to communicate erros via an explicit, seperate return value. This constrats with exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go's approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.
package main

import (
	"errors"
	"fmt"
)

// by convention, errors are the last return value and have type error, a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message
		return -1, errors.New("cant work with 42")
	}

	// a nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// it is possible to use custom types as errors by implementing the Error() method on them. Here's a variant of the example above that uses a custom type to explictly represenbt an arugment error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// in this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob
		return -1, &argError{arg, "cant work with it"}
	}

	return arg + 3, nil
}

func main() {
	// the two loops below test out each of our error-returning functions. note that the use of inline error check on the if line is a common idiom in Go code
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

// see this post for more on error handling
// https://go.dev/blog/error-handling-and-go
