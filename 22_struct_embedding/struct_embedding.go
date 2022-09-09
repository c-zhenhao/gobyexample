// Go supports embedding of structs and interfaces to express a more seamless composition of types. This is not to be confused with //go:embed which is ag o dreictive introduced in go 1.16 to embed files into the the application binary.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// a container embeds a base. an embedding looks like a field without a name
type container struct {
	base
	str string
}

func main() {
	// when creating structs with literals, we have to intialize the embedding explcitly; here the embedded types serves as a field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// we can access the base's fields directly on co, e.g. co.num
	fmt.Printf("c:={num: %v, str: %v}\n", co.num, co.str)

	// alternatively we can spell out the full path using the embdded type name
	fmt.Println("also num:", co.base.num)

	// since container embeds base, the methods of base are also become methods of a container.
	// Here we invoke a method that was embedded from base directly on co.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// embedding structs with methods may be used below to bestow interface implementations on other structs.
	// Here we see that a container now implements the describer interface because it embeds base
	var d describer = co
	fmt.Println("d.describe():", d.describe())
}
