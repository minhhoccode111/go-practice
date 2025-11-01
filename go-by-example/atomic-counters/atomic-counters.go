package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// multiple goroutines access the same variable for example
// func main() {
// 	var wg sync.WaitGroup
// 	ops := 0
// 	for range 50 {
// 		wg.Go(func() {
// 			for range 1000 {
// 				ops++
// 			}
// 		})
// 	}
// 	wg.Wait()
// 	fmt.Println(ops) // 42496
// }

// problem:
// use atomic when we want small and fast with a variable, like counter
// spawn 50 goroutines, each increase a variable 10000 times
// and wait for all with WaitGroup
func main() {
	var c atomic.Int64
	var wg sync.WaitGroup

	for range 10000 {
		wg.Go(func() {
			for range 100000 {
				c.Add(1)
			}
		})
	}

	wg.Wait()
	fmt.Println(c.Load())
}
