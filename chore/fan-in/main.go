package main

import (
	"fmt"
	"math/rand"
	"time"
)

// We can use a fan-in function to let whosoever ready talk
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for range 10 {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(s string) <-chan string { // return receive-only channel of string
	c := make(chan string)
	go func() { // we launch a goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return the channel to the caller
}
