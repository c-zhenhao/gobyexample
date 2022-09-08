// range iterates over elements in a variety of data structures.
// lets see how to use range with some of the data structures we already learned.

package main

import "fmt"

func main() {
	// here we use range to sum the numbers in a slice. arrays work liek this too
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// same as the above
	var sum2 int = 0
	for i := 0; i < len(nums); i++ {
		sum2 += nums[i]
	}
	fmt.Println("sum2:", sum2)

	// range on arrays and slices provides both the index and value for each entry. above we didnt need the index, so we ignored it with the blank identifier _. sometimes we actually want the index though
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// and the following for just the value
	for _, v := range kvs {
		fmt.Println("value:", v)
	}

	// range on strings iterates over Unicode code points. the first value is the starting byte index of the rune and the second the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
