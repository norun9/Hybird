package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"net/http"
	"time"
)

type IMessageController interface {
	List(c *gin.Context, p input.MessageList)
	Send(c *gin.Context, p input.MessageInput)
}

type messageController struct {
	messageInputBoundary usecase.IMessageInputBoundary
}

func NewMessageController(messageInputBoundary usecase.IMessageInputBoundary) IMessageController {
	return &messageController{messageInputBoundary}
}

func (mc *messageController) List(c *gin.Context, p input.MessageList) {
	ctx := c.Request.Context()
	messages, err := mc.messageInputBoundary.List(ctx, p)
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

func (mc *messageController) Send(c *gin.Context, p input.MessageInput) {
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
