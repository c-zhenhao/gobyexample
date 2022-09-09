// in the previous example we saw how to manage simple counter state using atomic operations
// for more complex state we can use a mutex to safely access data across multiple goroutines
// what is mutex?? -> mutual exclusive flag -- acts as a gatekeeper to a section of code allowing one thread in and blocking access to all others

package main

import (
	"fmt"
	"sync"
)

// container holds a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronise acces.
// note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// lock the mutex before accessing counters; unlock it at the end of function using a defer statement
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// note that the zero value of a mutex is usable as-is, so no intiliazation is required here
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// this function increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	wg.Wait() // wait for goroutines to finish
	fmt.Println(c.counters)
}

// running the program shows that counters updated as expected.
// next we'll look at implementing this same state managment task using only goroutines and channels
