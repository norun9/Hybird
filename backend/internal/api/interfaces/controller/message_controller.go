package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	ws "github.com/norun9/Hybird/pkg/websocket"
	"net/http"
)

type IMessageController interface {
	List(c *gin.Context, p input.MessageList)
	Send(c *gin.Context, _ interface{})
	Create(c *gin.Context, p input.MessageInput)
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

func (mc *messageController) Create(c *gin.Context, p input.MessageInput) {
	ctx := c.Request.Context()
	messages, err := mc.messageIB.Create(ctx, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (mc *messageController) Send(c *gin.Context, _ interface{}) {
	go ws.HandleMessages()
	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()
	ws.Clients[conn] = true
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			break
		}
		ws.Broadcast <- ws.Message{Type: t, Message: msg}
	}
}
