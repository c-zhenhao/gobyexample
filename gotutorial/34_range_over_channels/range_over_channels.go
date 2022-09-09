// in the previous example we saw how for and range provide iteration over basic data structures.
// we can also use this syntax to iterate over values recieved from a channel
package main

import "fmt"

func main() {

	// we'll iterate over 2 values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// the range iterates over each element as it's recieved from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements
	for elem := range queue {
		fmt.Println(elem)
	}
	fmt.Println(queue)
}

// this example also showed that its possible to close a non-empty channel but still have the remaining values be recieved.
