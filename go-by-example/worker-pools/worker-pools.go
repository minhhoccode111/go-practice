package main

import (
	"fmt"
	"time"
)

// problem:
// define channels to send jobs and receive results
// define worker function that takes in: id of the worker, a jobs channel to get jobs
// execute each job for 1 second, a results channel to send job's result to
// then in main thread, spawn 3 workers that ready to receive jobs
// send 5 jobs to jobs channel for those workers
// and then close the jobs channel, signaling that there are no more jobs
// read all the results from result channels (or use wait group)
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := range 3 {
		go worker(i, jobs, results)
	}

	for i := range numJobs {
		jobs <- i
	}
	close(jobs)

	for range numJobs {
		<-results
	}
}

func worker(id int, jobs chan int, results chan int) {
	for i := range jobs {
		fmt.Printf("worker: %d - job:    %d - start\n", id, i)
		time.Sleep(1 * time.Second)
		fmt.Printf("worker: %d - result: %d - finish\n", id, i*2)
		results <- i * 2
	}
}
