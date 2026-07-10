package main

/*
Select
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It
chooses one at random if multiple are ready.
*/

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan struct{})
	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		close(quit)
	}()
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

func fibonacci(c chan int, quit chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
