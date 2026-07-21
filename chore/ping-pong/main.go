package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {
	pings := make(chan string)
	pongs := make(chan string)
	go ping(pings, "tai vi sao")
	go pong(pings, pongs)
	fmt.Println(<-pongs)
}
