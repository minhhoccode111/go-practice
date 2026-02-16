package main

import (
	"fmt"
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
	fmt.Printf("Add a client: %s", ip)
}

func (h *hub) del(conn *websocket.Conn) {
	ip := conn.RemoteAddr().String()
	delete(h.clients, ip)
	fmt.Printf("Delete a client: %s", ip)
}

func (h *hub) broadcastMsg(text string) {
	fmt.Printf("Try broadcasting message to all clients: %s", text)
	for _, v := range h.clients {
		err := websocket.Message.Send(v, text)
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
	mux.Handle("/", websocket.Handler(handleWs))

	http.ListenAndServe(":9999", mux)
}

func handleWs(conn *websocket.Conn) {

}
