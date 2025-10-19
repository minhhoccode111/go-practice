package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pongs chan<- string, pings <-chan string) {
	pongs <- <-pings
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "passed message")
	go pong(pongs, pings)
	fmt.Println(<-pongs)
}
