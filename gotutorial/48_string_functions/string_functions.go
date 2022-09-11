// the standard library's strings package provides many useful string-related functions. Here are some examples to give you a sense of the package.
package main

import (
	"fmt"
	s "strings"
)

// we alias fmt.Println to a shorter name as we'll use it a lot below
var p = fmt.Println

// here's an example of the functions available in strings. since these are functions from the package, not methods on the string object itself, we need pass the string in question as the first argument to the function. you can find more functions in the strings package docs
func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))
}
