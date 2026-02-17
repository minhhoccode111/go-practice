package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	url := "ws://localhost:9999"
	conn, err := websocket.Dial(url, "", randomIP())
	if err != nil {
		log.Fatalf("Error occurs when dialing to %s: %s", url, err)
	}
	// what can i do with this client connection?
	// receive and print to stdout
	// read from stdin and send
	go listen(conn)
	send(conn)
}

func randomIP() string {
	return fmt.Sprintf(
		"http://%d.%d.%d.%d/",
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(255),
	)
}

func listen(conn *websocket.Conn) {
	var text string
	for {
		err := websocket.JSON.Receive(conn, &text)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("Server disconnected")
				break
			}
			log.Printf("Error occurs when receiving message: %s", err)
			continue
		}
		fmt.Println("Received from server: ", text)
	}
}

func send(conn *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		err := websocket.JSON.Send(conn, text)
		if err != nil {
			log.Printf("Error occurs when sending message: %s", err)
			continue
		}
		fmt.Println("Sent to server: ", text)
	}
}
