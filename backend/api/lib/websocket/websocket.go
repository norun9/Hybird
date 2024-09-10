package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	Ws *websocket.Conn
}

type Rooms struct {
	Clients []*Client
}

func (r *Rooms) AddClient(c *Client) {
	r.Clients = append(r.Clients, c)
}

func (r *Rooms) GetClients() []Client {
	var cs []Client
	for _, client := range r.Clients {
		cs = append(cs, *client)
	}
	return cs
}

func (r *Rooms) RemoveClient(client *Client) {
	for i, c := range r.Clients {
		if c == client {
			// Remove client2 from slice
			r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
			return
		}
	}
}

func (r *Rooms) Send(msg []byte) {
	for _, c := range r.Clients {
		err := c.send(msg)
		if err != nil {
			r.RemoveClient(c)
			c.Ws.Close()
		}
	}
}

func (c *Client) send(msg []byte) error {
	return c.Ws.WriteMessage(websocket.TextMessage, msg)
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// NOTE:to avoid the error: "message":"websocket: request origin not allowed by Upgrader.CheckOrigin"}
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
