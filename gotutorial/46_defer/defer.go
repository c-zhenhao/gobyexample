// defer is used to ensure that a function call is performed later in program's execution, usually for purposes of cleanup
// defer is often used where e.g. ensure and finally would be used in other languages
package main

import (
	"fmt"
	"os"
)

// suppose we wanted to create a file, write to it, and then close when we're done
// here's how we could do that with defer
func main() {
	// immediately after getting a file object with createFile, we defer the closing of that file with closeFile
	// this will be executed at the end of the enclosing function (main), after writeFile has finished
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// its important to check for errors when closing a file, even in a deferred function
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

// running the program confirms that the file is closed after being written
