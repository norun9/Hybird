package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/norun9/HyBird/backend/api/internal/usecase"
	"github.com/norun9/HyBird/backend/api/internal/usecase/dto/input"
	"github.com/norun9/HyBird/backend/api/internal/usecase/dto/output"
	myws "github.com/norun9/HyBird/backend/api/lib/websocket"
	"github.com/norun9/HyBird/backend/pkg/log"
	"go.uber.org/zap"
)

var rooms = &myws.Rooms{}

type IMessageController interface {
	List(c *gin.Context, p input.MessageList) ([]*output.MessageOutput, error)
	Receive(c *gin.Context, _ any) error
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

/*
*
@deprecated Transplant WebSocket function to Lambda
*/
func (mc *messageController) Receive(c *gin.Context, _ any) error {
	// TODO: コントローラーの責任が集中しているため、ビジネスロジックやデータ処理の部分を分離する
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
			log.Error("Failed to read message", zap.Error(err))
			break
		}
		rooms.Send(msg)
	}
	return nil
}
