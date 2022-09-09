// A Go string is a read-only slice of bytes. The language and standard library treat strings specially - as containers of text encoded in UTF-8. Strings are made of "characters". In Go, the concept of a character is called a rune - it's an integer that represents a Unicode code point.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s is a string assigned a literal value of "hello" in Thai language. Go strings literals are UTF-8 encoded
	const s = "สวัสดี"

	// since strings are equivalent to []byte, this will produce the length of the raw bytes stored within
	fmt.Println("len(s):", len(s)) // -> 18

	// indexing into a string produces hte raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	// to count how many runes are in a string, we can use the utf8 package. Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by multiple UTF-8 code point, so the result of this count may be surprising
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))
	fmt.Println("---")

	const r = "你好 世界"
	fmt.Println("len(r):", len(r)) // -> 15
	for i := 0; i < len(r); i++ {
		fmt.Printf("%x ", r[i])
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(r))
	fmt.Println("---")

	const t = "こんにちは世界"
	fmt.Println("len(t):", len(t)) // -> 15
	for i := 0; i < len(t); i++ {
		fmt.Printf("%x ", t[i])
	}
	fmt.Println("\nRune count:", utf8.RuneCountInString(t))
	fmt.Println("---")

	const h = "hello world"
	fmt.Println("len(h):", len(h)) // -> 11
	for i := 0; i < len(h); i++ {
		fmt.Printf("%x ", h[i])
	}
	fmt.Println("\nRune count:", utf8.RuneCountInString(h))
	fmt.Println("---")

	// a range loop handles strings specially and decodes each rune along with its offset in the string
	for index, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, index)
	}

	// we can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

// this demonstrates passing a rune value into a function
func examineRune(r rune) {

	// values enclosed in single quotes are rune literals. we can compare a rune value to a rune literal directly
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// !! double quotes - used to define a string. A string defined in double quotes will honour escaping charactes. For example, when print a string having \n there will be a line printerd. Similarly \t will have a tab printed

// back quotes - also used to define a string. A string encoded in backquotes is a raw ltieral string and doesn't honour any kind of escaping

// single quotes - to declare a byte or a rune we use single quotes. while declaring byte we have to specify the type. If we dont specify the type, then the default type is meant as a rune. A single quote will only allow one character. On declaring a byte or rune with two characters within a single quote, the compiler will rase an error -> 'invalid character literal (more than one character)'
