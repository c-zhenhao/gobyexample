// A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch. Here’s how to do it in Go.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Use time.Now with Unix, UnixMilli or UnixNano to get elapsed time since the Unix epoch in seconds, milliseconds or nanoseconds, respectively.
	now := time.Now()
	fmt.Println(now) // 2022-09-11 10:59:54.0965664 +0800 +08 m=+0.001615101

	fmt.Println(now.Unix())      // 1662865194
	fmt.Println(now.UnixMilli()) // 1662865194096
	fmt.Println(now.UnixNano())  // 1662865194096566400

	// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	fmt.Println(time.Unix(now.Unix(), 0))     // 2022-09-11 10:59:54 +0800 +08
	fmt.Println(time.Unix(0, now.UnixNano())) // 2022-09-11 10:59:54.0965664 +0800 +08
}
