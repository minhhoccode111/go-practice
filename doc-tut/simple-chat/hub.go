package main

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type hub struct {
	clients             map[string]*websocket.Conn
	addClientChannel    chan *websocket.Conn
	removeClientChannel chan *websocket.Conn
	broadcastChannel    chan string
}

func newHub() *hub {
	return &hub{
		clients:             make(map[string]*websocket.Conn),
		addClientChannel:    make(chan *websocket.Conn),
		removeClientChannel: make(chan *websocket.Conn),
		broadcastChannel:    make(chan string),
	}
}

func (h *hub) run() {
	for {
		select {
		case conn := <-h.addClientChannel:
			h.addClient(conn)
		case conn := <-h.removeClientChannel:
			h.removeClient(conn)
		case m := <-h.broadcastChannel:
			h.broadcast(m)
		}
	}
}

func (h *hub) addClient(conn *websocket.Conn) {
	h.clients[conn.RemoteAddr().String()] = conn
	fmt.Println("Clients <add>: ", h.clients)
}

func (h *hub) removeClient(conn *websocket.Conn) {
	delete(h.clients, conn.RemoteAddr().String())
	fmt.Println("Clients <del>: ", h.clients)
}

func (h *hub) broadcast(m string) {
	fmt.Println("broadcast message: ", m)
	for _, conn := range h.clients {
		err := websocket.JSON.Send(conn, m)
		if err != nil {
			fmt.Println("error in Broadcast message: ", err.Error())
			continue
		}
	}
}
