// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

package main

import "fmt"

// unlike arrays, slices are typed only by the elements they contain (not the number of elements).
func main() {
	// to create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initally zero-valued).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// we can set and get just like with arrays
	s[0] = "a"
	fmt.Println(s) // -> [a  ]
	s[1] = "b"
	fmt.Println(s) // -> [a b ]
	s[2] = "c"
	fmt.Println(s) // -> [a b c]

	// panic: runtime error: index out of range [3] with length 3
	// s[3] = "d"
	// fmt.Println(s)

	fmt.Println("set:", s)    // -> set: [a b c]
	fmt.Println("get:", s[2]) // -> get: c

	// to get the length
	fmt.Println("len:", len(s)) // -> len: 3

	// in addition to basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")      // -> [a b c d]
	s = append(s, "e", "f") // -> [a b c d e f]
	fmt.Println("apd:", s)  // -> apd: [a b c d e f]

	// slices can also be copied.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // -> cpy: [a b c d e f] same as s

	// slices support a "slice" operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]            // 2 <= index < 5, i.e. 5 is exclusive
	fmt.Println("sl1:", l) // -> sl1: [c d e]

	l2 := s[:5]             // 0, 1, 2, 3, 4 (5 is exclusive)
	fmt.Println("sl2:", l2) // -> sl2: [a b c d e]

	l3 := s[2:]             // 2, 3, 4, 5 (6 is exclusive)
	fmt.Println("sl3:", l3) // -> sl3: [c d e f]

	// golang doesnt have steps in slice (i.e. [2:5:7] is not supported)
	// Slices in go are continuous in memory, so we can use the following to create a slice from an array
	a := [6]string{"a", "b", "c", "d", "e", "f"}
	numbers := [5]int{1, 2, 3, 4, 5}
	// a[2:6:2]
	for i := 2; i < 6; i += 2 {
		fmt.Println(a[i])
		fmt.Println(numbers[i])
	}

	// to revisit next time -> whats the point of 3rd index in slice
	// https://www.ardanlabs.com/blog/2013/12/three-index-slices-in-go-12.html

	// we can declare and init a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

	// Note that while slices are different types than arrays, they are rendered similarly by fmt.Println.
}
