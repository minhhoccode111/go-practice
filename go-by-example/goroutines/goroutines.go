package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	/*
	   $ go run goroutines.go
	   direct : 0
	   direct : 1
	   direct : 2
	   going
	   goroutine : 0
	   goroutine : 1
	   goroutine : 2
	   done
	*/
}
