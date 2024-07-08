package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"github.com/norun9/Hybird/internal/api/usecase/dto/output"
	ws "github.com/norun9/Hybird/pkg/websocket"
)

type IMessageController interface {
	List(c *gin.Context, p input.MessageList) ([]*output.MessageOutput, error)
	Send(c *gin.Context, _ interface{}) error
	Create(c *gin.Context, p input.MessageInput) (*output.MessageOutput, error)
}

type messageController struct {
	messageIB usecase.IMessageInputBoundary
}

func NewMessageController(messageIB usecase.IMessageInputBoundary) IMessageController {
	return &messageController{
		messageIB,
	}
}

func (mc *messageController) List(c *gin.Context, p input.MessageList) ([]*output.MessageOutput, error) {
	ctx := c.Request.Context()
	return mc.messageIB.List(ctx, p)
}

func (mc *messageController) Create(c *gin.Context, p input.MessageInput) (*output.MessageOutput, error) {
	ctx := c.Request.Context()
	return mc.messageIB.Create(ctx, p)
}

func (mc *messageController) Send(c *gin.Context, _ interface{}) error {
	go ws.HandleMessages()
	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	ws.Clients[conn] = true
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		ws.Broadcast <- ws.Message{Type: t, Message: msg}
	}
	return nil
}
