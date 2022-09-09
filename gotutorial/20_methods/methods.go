// Go supports methods defined on struct types

package main

import "fmt"

type rect struct {
	width, height int
}

// this area method has a reciever type of *rect
// *rect <-- i.e. takes a rect pointer
func (r *rect) area() int {
	return r.width * r.height
}

// methods can defined for either pointer or value reciever types. Here's an example of a value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// here we call the 2 methods defined for our struct
func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving
	rp := &r  // &r syntax gives the memory address of r -> 0xc00007fea0area
	print(&r) // 0xc00007fea0area: 50
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
