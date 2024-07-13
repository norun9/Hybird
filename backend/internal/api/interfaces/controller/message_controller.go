package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/Hybird/internal/api/usecase"
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"github.com/norun9/Hybird/internal/api/usecase/dto/output"
	"github.com/norun9/Hybird/pkg/log"
	myws "github.com/norun9/Hybird/pkg/websocket"
	"go.uber.org/zap"
)

var rooms = &myws.Rooms{}

type IMessageController interface {
	List(c *gin.Context, p input.MessageList) ([]*output.MessageOutput, error)
	Receive(c *gin.Context, _ interface{}) error
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

func (mc *messageController) Receive(c *gin.Context, _ interface{}) error {
	conn, err := myws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := &myws.Client{Ws: conn}
	rooms.AddClient(client)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Logger.Error("Failed to read message", zap.Error(err))
			break
		}
		rooms.Send(msg)
	}
	return nil
}
