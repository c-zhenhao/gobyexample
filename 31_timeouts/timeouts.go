// timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
// implementing timeouts in Go is easy and elegant thanks to channels and select
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	// for our example, suppose we're executing an extenral call that return its result on a channel c1 after 2s.
	// note that the channel is buffered, so the send in the goroutine is nonblocking. this is a common pattern to prevent goroutine leaks in case the channel is never read.
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// here's the select implementing a timeout. res := <-c1 awaits the result and the <-time.After awaits a value to be sent after the timeout of 1s. since select proceeds with the first receive that's ready, we'll take the timeout case if the operation takes more than the allowed 1s.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// running this program shows the first operation timing out and the second succeeding.
