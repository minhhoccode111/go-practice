package main

import (
	"fmt"
	"sync"
	"time"
)

// problem:
// define a wait group
// define worker that has an id and run task in 1 second
// spawn 5 goroutines for 5 workers
// and wait for all to finish
func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Go(func() {
			worker(i)
		})
	}

	wg.Wait()
}

func worker(id int) {
	fmt.Println("worker start", id)
	time.Sleep(1 * time.Second)
	fmt.Println("worker end", id)
}
