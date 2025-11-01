package main

import (
	"fmt"
	"time"
)

// problem:
// normal rate-limiter: loop through ticker channel and handle jobs in another channel each time ticker tick
// bursty rate-limiter: the same idea as above but the channel start with 3 existing items in the ticker channel
func main() {
	requests := make(chan int, 5)
	for i := range 5 {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(200 * time.Millisecond)
	for req := range requests {
		t := <-limiter.C
		fmt.Println("request:", req, "time:", t)
	}

	burstyLimiter := make(chan time.Time, 3)
	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := range 5 {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		t := <-burstyLimiter
		fmt.Println("request:", req, "time:", t)
	}
}
