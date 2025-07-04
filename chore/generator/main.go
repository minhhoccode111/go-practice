package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GENERATOR: FUNCTION THAT RETURNS A CHANNEL
func main() {
	c := boring("boring!") // function returning a channel
	for range 5 {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func main0() {
	joe := boring("Joe")
	ann := boring("Ann")
	for range 5 {
		// have to wait for each other, so we need fan-in to make them execute
		// independent
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're boring; I'm leaving.")
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
