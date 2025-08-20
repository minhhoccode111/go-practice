package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nStr := os.Args[1]
	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		fmt.Print("Fatal", "\n")
		return
	}

	sieve(n)
}

// The prime sieve: Daisy-chain filter  processes together.
func sieve(n int64) {
	ch := make(chan int) // create a new channel
	var count int64 = 0
	go generate(ch) // start generate() as a subprocess
	for {
		prime := <-ch
		count++
		if count == n {
			fmt.Print(prime, "\n")
			return
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // send 'i' to channel 'dst'.
		}
	}
}

// send the sequence 2, 3, 4, ... to channel 'ch'
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}
