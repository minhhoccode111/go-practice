package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanInSelect(boring("Joe"), boring("Ann"))
	for range 10 {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

// Similar to original Fan-In function but use select
func fanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
