package main

import (
	"fmt"
)

// problem: create jobs, done channels
// spawn a goroutine to received and print the value until the jobs channel is closed
// on the main thread, send 3 values to jobs channel then close it
// print the last time to confirm there is no more job
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool) // should this be buffered?

	go func() {
		for j := range jobs {
			fmt.Println("received job: ", j)
		}
		fmt.Println("received all jobs")
		done <- true
	}()

	for i := range 3 {
		fmt.Println("sent job:", i)
		jobs <- i
	}
	close(jobs)
	fmt.Println("send all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more job", ok)
}
