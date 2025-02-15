package main

/*
sync.Mutex
We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only
one goroutine can access a variable at a time to avoid conflicts?

This concept is called mutual exclusion, and the conventional name for the data
structure that provides it is mutex.

Go's standard library provides mutual exclusion with sync.Mutex and its two
methods:

	Lock
	Unlock

We can define a block of code to be executed in mutual exclusion by surrounding
it with a call to Lock and Unlock as shown on the Inc method.

We can also use defer to ensure the mutex will be unlocked as in the Value
method.
*/

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex     // mutex to synchronize access to the map
	v  map[string]int // map to store key-value pairs
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	(*c).mu.Lock()         // lock the mutex to ensure exclusive access
	defer (*c).mu.Unlock() // ensure the mutex is unlocked after modification
	(*c).v[key]++          // increase the value for the key
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	(*c).mu.Lock()         // lock the mutex before reading
	defer (*c).mu.Unlock() // ensure the mutex is unlocked when the function returns
	return (*c).v[key]     // return the value for the key
}

func main() {
	c := SafeCounter{v: make(map[string]int)} // initialize SafeCounter with an empty map
	for i := 0; i < 1000; i++ {               // start 1000 goroutines
		go c.Inc("somekey") // to increment "somekey" concurrently
	}
	time.Sleep(time.Second)         // wait for goroutines to finish
	fmt.Println(c.Value("somekey")) // 1000
}
