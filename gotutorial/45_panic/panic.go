// a panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn't occur during normal operation, or that we aren't prepared to handle gracefully
package main

import "os"

func main() {
	// we'll use panic throughout this site to check for unexpected errors. This is the only program on the site designed to panic
	// panic("a problem")

	// a common use of panic is to abort if a function returns an error value that we don't know how to (or want to) handle. Here's an example of panicking if we get an unexpected error when creating a new file
	_, err := os.Create("./file")
	if err != nil {
		panic(err)
	}
}

// running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status

// when first panic in main fires, the program exits without reaching the rest of the rest of the code. if you'd like to see the program try to create a temp file, comment that first panic out

// note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible
