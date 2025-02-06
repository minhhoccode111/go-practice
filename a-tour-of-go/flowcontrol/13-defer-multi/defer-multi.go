package main

import "fmt"

/*
Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its
deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this blog post.
*/

func main() {
	printMultiDefer()
}

func printMultiDefer() {
	fmt.Println("Start counting")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Done!")

	/*
	   counting
	   done
	   9
	   8
	   7
	   6
	   5
	   4
	   3
	   2
	   1
	   0
	*/
}
