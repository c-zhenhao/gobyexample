// Go's structs are typed collections of fields. They are useful for grouping data together to form records

package main

import "fmt"

// This person struct type has name and age fields
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// you can safely return a pointer to local variable as a local variable will survive the scope of the function
	return &p
}

func main() {
	// this syntax creates a new struct
	fmt.Println(person{"Bob", 20}) // -> {Bob 20}

	// you can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30}) // -> {Alice 30}

	// omitted fields will be zero valued
	fmt.Println(person{name: "Fred"}) // -> {Fred 0}

	// an & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40}) // -> &{Ann 40}

	// it's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon")) // -> {Jon 42}

	// access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // -> Sean
	fmt.Println(s.age)  // -> 50

	// you can also use dots with struct pointers - the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age) // -> 50

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age) // -> 51
}
