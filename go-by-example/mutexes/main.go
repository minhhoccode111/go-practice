package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to update it concurrently
// from multiple goroutines, we add a Mutex to synchronize access. Note that
// mutexes must not be copied, so if this struct is passed around, it should be
// done by pointer
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing counter; unlock it at the end of the function
// using a defer statement
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// Note that the zero value of a mutex is usable as-is, so no initialization
	// is required here
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// this function increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently; note that they all access the same
	// Container, and two of them access the same counter
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}

/*
Running the program shows that the counters updated as expected
$ go run main.go
map[a:20000 b:10000]
*/
