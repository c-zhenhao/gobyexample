// the primary mechanism for managing state in Go is communication over channels.
// we saw this for example with worker pools
// there are a few other options for managing state though. here we'll look at using the sync/atomic package for atomic counters accessed by multiple goroutines
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// difference between signed and unsigned integers
// https://medium.com/geekculture/signed-and-unsigned-integers-b5d4e7f0369f
// The “u” in front of each variant means unsigned, meaning that they can only hold positive values and the variants without the “u” are signed, meaning that can hold both negative and positive values.

func main() {
	// we'll use an unsigned integer to represent our (always-positive) counter
	var ops uint64

	// a WaitGroup will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// We'll start 50 goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// wait util all goroutines are done
	wg.Wait()

	// it is safe to access ops now because we know no other goroutine is writing to it. reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64
	fmt.Println("ops:", ops)
}

// we expect to get exactly 50,000 operations. had we used a non-atomic ops++ to increment the counter, we'd likely get a different number, changing between runs, because the goroutines would interfere with each other. Moreover, we'd get data race failures when running with the -race flag
