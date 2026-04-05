package main

import (
	"fmt"
	"sync"
	"time"
)

func startWorker(worker int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for i := range jobs {
		fmt.Printf("worker %d starts job %d\n", worker+1, i+1)
		time.Sleep(1 * time.Second)
		// fan-in jobs
		results <- i
		fmt.Printf("worker %d finishes job %d\n", worker+1, i+1)
	}
	wg.Done()
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)
		go startWorker(i, jobs, results, &wg)
	}

	// fan-out jobs
	go func() {
		for i := range numJobs {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := range results {
		fmt.Printf("results receive job %d\n", i+1)
	}
}
