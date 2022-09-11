// Go makes it possible to recover from a panic, by using the recover built-in function. A recover can stop a panic from aborting the program and let it continue with execution instead.
// an example of where this can be useful: a server wouldn't want to crash if a panic occurs, but instead would want to recover from the panic and continue to serve requests. in fact, this is what Go's net/http does by default for HTTP servers
package main

import "fmt"

// this function panics
func mayPanic() {
	panic("a problem")
}

func main() {

	// recover must be called within a deferred function. when the enclosing function panics, the defer will activate and a recover call within it will catch the panic.

	defer func() {

		// the return value of recover is the error raised in the call to panic
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// this code will not run, because mayPanic panics. The execution of main stops at the point of the panic and reumes in the deferred closure
	fmt.Println("After mayPanic()")
}
