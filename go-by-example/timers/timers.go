package main

import (
	"fmt"
	"time"
)

// problem:
// create a timer that will fire after 2 seconds
// wait (block main thread) for the timer channel to fire
// create a timer that will fire after 1 second
// wait (in other goroutine) for the timer channel to fire
// stop the timer before it has chance to fire
// literally wait for 2 more seconds to see that it has enough time to fire but wont
func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped")
	}
	/*
	   $ go run timers.go
	   timer1 fired
	   timer2 stopped
	*/
}
