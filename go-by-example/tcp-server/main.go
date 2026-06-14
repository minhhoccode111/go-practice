package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("error listening:", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection: %s", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	msg, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("error reading connection: %s", err)
		return
	}
	ackMsg := fmt.Sprintf("ACK: %s\n", msg)
	response := strings.ToUpper(strings.TrimSpace(ackMsg))
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("error writing response: %s", err)
	}
}
