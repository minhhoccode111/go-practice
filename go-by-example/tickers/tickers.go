package main

import (
	"fmt"
	"time"
)

// problem:
// define a ticker that tick every 500ms
// listen for its tick on a separate goroutine
// then stop after 1600ms
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done tickering")
				return
			case t := <-ticker.C:
				fmt.Println("tick at:", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	close(done)
	time.Sleep(10 * time.Millisecond) // to have time to print 'done tickering'
	fmt.Println("ticker stopped")
}
