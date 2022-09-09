// to wait for multiple goroutines to finish, we can use a wait group
package main

import (
	"fmt"
	"sync"
	"time"
)

// this is a function we'll run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// this wait group is used to wait for all goroutines launched here to finish
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer
	var wg sync.WaitGroup

	// launch several goroutines and increment the wait group counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// avoid re-use of the same i value in each goroutine closure.
		// see https://go.dev/doc/faq#closures_and_goroutines
		i := i

		// wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// block until the wait group counter goes back to 0; all the workers notified they're done
	wg.Wait()
}

// Note that this approach has no straightforward way to propogate errors from workers. For more advanced use cases, consider using the errgroup package
// The order of workers starting up and finishing is likely to be different for each invocation
