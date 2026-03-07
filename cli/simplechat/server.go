package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type hub struct {
	clients   map[string]*websocket.Conn
	addCh     chan *websocket.Conn
	delCh     chan *websocket.Conn
	broadcast chan string
}

func newHub() *hub {
	return &hub{
		clients:   make(map[string]*websocket.Conn),
		addCh:     make(chan *websocket.Conn),
		delCh:     make(chan *websocket.Conn),
		broadcast: make(chan string),
	}
}

func (h *hub) run() {
	for {
		select {
		case conn := <-h.addCh:
			h.add(conn)
		case conn := <-h.delCh:
			h.del(conn)
		case text := <-h.broadcast:
			h.broadcastMsg(text)
		}
	}
}

func (h *hub) add(conn *websocket.Conn) {
	ip := conn.RemoteAddr().String()
	h.clients[ip] = conn
	log.Printf("Add a client: %s", ip)
	log.Printf("Current clients: %v", h.clients)
}

func (h *hub) del(conn *websocket.Conn) {
	ip := conn.RemoteAddr().String()
	delete(h.clients, ip)
	log.Printf("Delete a client: %s", ip)
}

func (h *hub) broadcastMsg(text string) {
	log.Printf("Broadcasting message [%s] to all clients", text)
	for _, v := range h.clients {
		err := websocket.JSON.Send(v, text)
		if err != nil {
			fmt.Printf("Error occurs when sending message to %s: %s", v.RemoteAddr().String(), err)
			continue
		}
	}
}

func main() {
	h := newHub()
	go h.run()

	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(c *websocket.Conn) {
		handleWs(c, h)
	}))

	http.ListenAndServe(":9999", mux)
}

func handleWs(conn *websocket.Conn, h *hub) {
	h.addCh <- conn
	var text string
	for {
		err := websocket.JSON.Receive(conn, &text)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("Client %s disconnected", conn.RemoteAddr().String())
			} else {
				log.Printf("Error occurs when receiving message: %s", err)
			}
			h.delCh <- conn
			break
		}
		h.broadcast <- text
	}
}
