package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeoutSelect()
	timeoutWhole()
}

func timeoutSelect() {
	c := boring("Joe")
	for range 10 {
		select {
		case s := <-c:
			fmt.Println(s)
		// the time.After function returns a channel that blocks for the
		// specified duration. After the interval, the channel delivers the
		// current time, once.
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}

// timeout for while conversation using select
func timeoutWhole() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for range 10 {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}

func boring(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1100)) * time.Millisecond)
		}
	}()
	return c
}
