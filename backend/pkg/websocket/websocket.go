package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type Message struct {
	Type    int
	Message []byte
}

var Broadcast = make(chan Message)

var Clients = make(map[*websocket.Conn]bool)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleMessages() {
	for {
		message := <-Broadcast
		for client := range Clients {
			err := client.WriteMessage(message.Type, message.Message)
			if err != nil {
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
