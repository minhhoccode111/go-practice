package main

import (
	"fmt"
	"sync"
	"time"
)

// practice implement worker pool in Go
// 3 workers
// a channel can store at most 5 jobs
// 10 jobs will be sent to that channel
// each job takes 1 second to process

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, 5)

	go func() {
		for i := range numJobs {
			jobs <- i
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for i := range numWorkers {
		wg.Add(1)
		go startWorker(i, jobs, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}

func startWorker(worker int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range jobs {
		process(worker, i)
	}
}

func process(worker, job int) {
	fmt.Printf("worker %d started processing job %d\n", worker+1, job+1)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d finished processing job %d\n", worker+1, job+1)
}
