package main

/*

Default Selection
The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:

select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
*/

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(" .")
			time.Sleep(200 * time.Millisecond)
		}
	}
	/*
		output:
		.
		.
		tick.
		.
		.
		tick.
		.
		.
		tick.
		.
		.
		tick.
		.
		.
		BOOM!
	*/
}
