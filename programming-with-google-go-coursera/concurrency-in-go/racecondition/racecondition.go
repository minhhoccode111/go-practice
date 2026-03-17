package main

import (
	"fmt"
	"time"
)

var global = 0

func main() {
	// goroutine 1
	go func() {
		for range 1_000_000_000 {
			global++
		}
	}()

	// goroutine 2
	go func() {
		for range 1_000_000_000 {
			global--
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(global)

	/*
		A race condition exists when two or more operations access the same
		memory location concurrently, at least one operation is a write, and
		there is no synchronization guaranteeing ordering.

		$ go run .
		4003793
	*/
}
