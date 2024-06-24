package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/norun9/Hybird/internal/api/usecase"
	"net/http"
	"time"
)

type MessageController struct {
	interactor usecase.IMessageInputBoundary
}

func NewMessageController(interactor usecase.IMessageInputBoundary) *MessageController {
	return &MessageController{interactor}
}

func (mc *MessageController) List(c *gin.Context) {
	messages, err := mc.interactor.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (mc *MessageController) Send(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		time.Sleep(time.Second)
	}
}
