package main

import (
	"fmt"
	"math/rand"
	"time"
)

// We can use a fan-in function to let whosoever ready talk
func main() {
	// create a channel from 2 input channels
	c := fanIn(boring("Joe"), boring("Ann"))
	// only read 10 results
	for range 10 {
		fmt.Println(<-c)
	}
	// exit program
	fmt.Println("You're both boring; I'm leaving.")
}

// fanIn is a function that
// receive 2 send-only string channels
// return a send-only string channel
func fanIn(input1, input2 <-chan string) <-chan string {
	// create a channel of string
	c := make(chan string)
	// spawn first goroutine so that the main execution continue without waiting
	go func() {
		// loop forever
		for {
			// forward anything from the input channel to output channel
			c <- <-input1
		}
	}()
	// spawn second goroutine so that the main execution continue without waiting
	go func() {
		// loop forever
		for {
			// forward anything from the input channel to output channel
			c <- <-input2
		}
	}()
	// return channel
	return c
}

// boring is a function that
// receive a string
// return a send-only string channel
func boring(s string) <-chan string {
	// create a channel of string
	c := make(chan string)
	// spawn a goroutine so that the main execution continue without waiting
	go func() {
		// loop forever
		for i := 0; ; i++ {
			// and send a string to channel
			c <- fmt.Sprintf("%s %d", s, i)
			// then sleep a random time
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	// return the channel to the caller
	return c
}
