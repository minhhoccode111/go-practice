package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

/*
=== PSEUDO CODE ===
1. write hub.go
create a hub struct group those data:
- clients: a map of all clients - websocket connections
- addClientChannel: channel to add client - send a connection to this channel will add it to the clients map
- removeClientChannel: channel to remove client - send a connection to this channel will remove it from the clients map
- broadcastChannel: channel to broadcast message - send a message to this channel will broadcast it to all clients
a newHub function that returns a new default hub
a hub's run method to start listening for all channels
a hub's addClient method to add a client to the clients map
a hub's removeClient method to remove a client from the clients map
a hub's broadcast method to broadcast/send a message to all clients
2. write main.go
init newHub and run in separate goroutine
create a mux, register a websocket handler, and serve the server
the websocket handler:
- send the connection to the addClientChannel of the hub
- then continuously wait for messages from the connection and send them to the broadcastChannel, which will then be broadcast to all clients.
  Only stop when there is an error or the connection is closed
3. write client.go
create a websocket connection to the server using a fake IP address
run a goroutine that continuously wait for messages from the connection and print them to stdout. Only stop when there is an error
read any stdin and send it through the connection to the server
*/

func main() {
	h := newHub()
	go h.run()
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(c *websocket.Conn) {
		wsHandler(c, h)
	}))
	server := http.Server{
		Addr:    ":9999",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func wsHandler(conn *websocket.Conn, h *hub) {
	h.addClientChannel <- conn
	for {
		var m string
		err := websocket.JSON.Receive(conn, &m)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("client disconnected: ", conn.RemoteAddr().String())
			} else {
				fmt.Println("error in Receive Message: ", err.Error())
			}
			h.removeClientChannel <- conn
			break
		}
		h.broadcastChannel <- m
	}
}
