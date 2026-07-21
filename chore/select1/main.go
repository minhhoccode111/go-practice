package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "c2"
	}()

	for {
		select {
		case s := <-c1:
			fmt.Println(s)
		case s := <-c2:
			fmt.Println(s)
		case <-time.After(3 * time.Second):
			fmt.Println("program finished")
			return
		}
	}
}
