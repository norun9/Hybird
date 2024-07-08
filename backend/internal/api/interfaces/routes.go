package interfaces

import (
	"github.com/norun9/Hybird/internal/api/interfaces/controller"
)

type Method string

var (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Delete Method = "DELETE"
)

type Path struct {
	Path string
	Method
}

type Handler struct {
	Func interface{}
}

func GetMapRoute(messageController controller.IMessageController) map[Path]Handler {
	return map[Path]Handler{
		{"/v1/messages", Get}:    {messageController.List},
		{"/v1/messages/ws", Get}: {messageController.Send},
		{"/v1/messages", Post}:   {messageController.Create},
	}
}
