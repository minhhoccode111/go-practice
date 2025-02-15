package main

/*
Select
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It
chooses one at random if multiple are ready.
*/

import "fmt"

func main() {
	// make a channel `c` to send and receive int from goroutine Fibonacci
	c := make(chan int)
	// make a channel `quit` to send signal when to stop the program
	quit := make(chan int)

	// anonymous goroutine to receive and print out 10 numbers of Fibonacci
	// from channel `c`
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // receive value from channel `c` and print out
		}
		quit <- 0 // after receiving 10 numbers, send a signal to channel `quit`
	}()

	// call func `fibonacci` to initialize and send fibonacci sequence to
	// channel `c`
	fibonacci(c, quit)
	/*
		0
		1
		1
		2
		3
		5
		8
		13
		21
		34
		quit
	*/
}

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1 // init 2 starting numbers of the fibonacci sequence
	for {
		select {
		case c <- x: // send current fibonacci number to channel `c`
			x, y = y, x+y // update the next fibonacci number
		case <-quit: // when receive signal from channel `quit` stop the program
			fmt.Println("quit")
			return
		}
	}
}
