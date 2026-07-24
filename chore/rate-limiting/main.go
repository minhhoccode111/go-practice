package main

import (
	"fmt"
	"time"
)

func main() {
	const jobs = 5

	// without burst
	fmt.Println("without burst")
	requests := make(chan int, jobs)
	go func() {
		for i := range jobs {
			requests <- i
		}
		close(requests)
	}()

	limiter := time.Tick(500 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request:", req, "time:", time.Now())
	}

	// with burst
	fmt.Println("with burst")
	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, jobs)
	for i := range jobs {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request:", req, "time:", time.Now())
	}
}
