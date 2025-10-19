package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case message := <-messages:
		fmt.Println(message)
	default:
		fmt.Println("no message received")
	}

	message := "hi"
	select {
	case messages <- message:
		fmt.Println(message)
	default:
		fmt.Println("no message sent")
	}

	select {
	case message := <-messages:
		fmt.Println(message)
	case signal := <-signals:
		fmt.Println(signal)
	default:
		fmt.Println("no activity")
	}
}
