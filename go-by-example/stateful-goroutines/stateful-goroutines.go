package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
problem:
- communication by goroutines and channels
- readOp struct to represent a read operation, contain a key we want to read from map and a resp chan to receive value back
- writeOp struct to represent a write operation, contain a key-value we want to set to map and a resp chan to receive signal back indicate that operation succeed
- create 2 uint64s to count operations
- create 2 channels to receive read and write requests
- spawn a goroutine act as server that have a map "database" and a "for" + "select" combination to handle requests send to reads and writes channel
- spawn 100 goroutines that continuously send read requests to "server" forever with random key, wait for response, increase count by one, and sleep one millisecond
- spawn 10 goroutines that continuously send write requests to "server" forever with random key-value, wait for response, increase count by one, and sleep one millisecond
- sleep one second for program to finish
- then load and print total read and write operations
*/

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)
	done := make(chan bool)

	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			case <-done:
				return
			}
		}
	}()

	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				fmt.Println(<-read.resp)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	close(done)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("read ops:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("write ops:", writeOpsFinal)
}
