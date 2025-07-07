package main

import "fmt"

// function 'f' takes two channels of type int
func f(left, right chan int) {
	// 'left' will have to wait until 'right' send something + 1
	left <- 1 + <-right
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	// left and right point to the same leftmost when init
	left := leftmost
	right := leftmost
	// loop 100000 times
	for range n {
		// assign right to be a new channel
		right = make(chan int)
		// make left wait until receive something from right, and plus 1
		go f(left, right)
		// assign right to left and continue the loop til the end
		left = right
	}
	// send '1' to the right channel, which keep passing the number 1 through
	// every channel and increase by 1, until we reach leftmost
	go func(c chan int) { c <- 1 }(right)
	// print leftmost
	fmt.Println(<-leftmost) // 100001
}
