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

// func main() {
// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)
// 	for {
// 		select {
// 		case <-tick:
// 			fmt.Println("tick.")
// 		case <-boom:
// 			fmt.Println("BOOM!")
// 			return
// 		default:
// 			fmt.Println("    .")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// }

func main() {
	// create a channel `tick` to send signal every 100 milliseconds
	tick := time.Tick(100 * time.Millisecond)
	// create a channel `boom` to send only one signal after 500 milliseconds
	boom := time.After(500 * time.Millisecond)

	// infinite loop to handle signals from `tick` and `boom`
	for {
		select {
		case <-tick: // when receive signal from `tick`
			fmt.Println("tick.") // print "tick."
		case <-boom: // when receive signal from `boom`
			fmt.Println("BOOM!") // print "BOOM!"
			return               // and exit program
		default: // if there is no signal from `tick` or `boom`
			fmt.Println("     .")             // print "    ."
			time.Sleep(50 * time.Millisecond) // then sleep for 50 milliseconds
			// to prevent the loop from running too fast
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
