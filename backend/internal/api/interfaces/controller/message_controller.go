package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"log"
	"net/http"
)

type Message struct {
	Type    int
	Message []byte
}

var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type IMessageController interface {
	List(c *gin.Context, p input.MessageList)
	Send(c *gin.Context, p input.MessageInput)
}

type messageController struct {
	messageIB usecase.IMessageInputBoundary
}

func NewMessageController(messageIB usecase.IMessageInputBoundary) IMessageController {
	return &messageController{
		messageIB,
	}
}

func (mc *messageController) List(c *gin.Context, p input.MessageList) {
	ctx := c.Request.Context()
	messages, err := mc.messageIB.List(ctx, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (mc *messageController) Send(c *gin.Context, p input.MessageInput) {
	// IBで行う
	go handleMessages()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()
	clients[conn] = true
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			break
		}
		// 受け取ったメッセージをbroadcastを通じてhandleMessages()関数へ渡す
		broadcast <- Message{Type: t, Message: msg}
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteMessage(message.Type, message.Message)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
