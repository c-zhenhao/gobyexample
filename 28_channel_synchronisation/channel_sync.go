// we can use channels to synchronise execution across goroutines
// here's an example of using a blocking recieve to wait for a goroutine to finish. when waiting for multiple goroutines to finish, you may prefer to use a WaitGroup
package main

import (
	"fmt"
	"time"
)

// this is the function we will use in a goroutine. the done channel will be used to notify another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value to notify we are done
	done <- true
}

func main() {
	done := make(chan bool, 1)
	// start a worker gorountine, giving it the channel to notify on
	go worker(done)

	// block until we receive a notification from the worker on the channel
	<-done // if you removed the <- done line from this program, the program would exit before the worker even started
}
