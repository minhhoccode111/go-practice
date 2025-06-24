package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id     int
	digits int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// calculate sum of digits of a number and sleep for 2 seconds
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	// each worker will execute all jobs in jobs channel and send result to results channel
	for job := range jobs {
		output := Result{job, digits(job.digits)}
		results <- output
	}
	// then call done to alert WaitGroup
	wg.Done()
}

func createWorkerPool(noOfWorker int) {
	// wait group to wait for each worker
	var wg sync.WaitGroup
	// spawn worker base on noOfWorker
	for range noOfWorker {
		// increase wait group
		wg.Add(1)
		// each worker will do all current jobs in channel
		go worker(&wg)
	}
	// wait for all workers to finish
	wg.Wait()
	// close results channel
	close(results)
}

func allocate(noOfJobs int) {
	// send jobs to jobs channel a number of time
	for i := range noOfJobs {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.digits, result.sumOfDigits)
	}
	// send done signal
	done <- true
}

func main() {
	// mark start time
	startTime := time.Now()
	// number of jobs to will create
	noOfJobs := 100
	// create number of jobs concurrently and send to jobs channel, then close the channel later
	// in this case channel has capacity of 10, so channel will block until a
	// worker come and pick up 10 jobs in the channel
	go allocate(noOfJobs)
	// a channel to indicate done (instead of using sync.WaitGroup)
	done := make(chan bool)
	// handle every item send to resutls channel, then send a signal to done channel
	go result(done)
	// number of workers to run jobs
	noOfWorkers := 100
	// spawn a number of workers, each worker will handle every jobs in jobs channel
	createWorkerPool(noOfWorkers)
	<-done
	// mark end time
	endTime := time.Now()
	// diff between start time and end time
	diff := endTime.Sub(startTime)
	// print result
	fmt.Println("Total time take: ", diff.Seconds(), "seconds")
}
