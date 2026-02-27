package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	conn, err := websocket.Dial("ws://localhost:9999", "", CreateDemoIP())
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	go receive(conn)
	send(conn)
}

func CreateDemoIP() string {
	const l = 4
	var array [l]int
	for i := range l {
		array[i] = rand.Intn(256)
	}
	return fmt.Sprintf(
		"http://%d.%d.%d.%d",
		array[0], array[1], array[2], array[3],
	)
}

func receive(conn *websocket.Conn) {
	for {
		var m string
		err := websocket.JSON.Receive(conn, &m)
		if err != nil {
			log.Println("Error in Receive Message: ", err.Error())
			continue
		}
		fmt.Println("Message from Server: ", m)
	}
}

func send(conn *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Send message to Server: ", text)
		err := websocket.JSON.Send(conn, text)
		if err != nil {
			fmt.Println("Error in Send Data: ", err.Error())
			continue
		}
	}
}
