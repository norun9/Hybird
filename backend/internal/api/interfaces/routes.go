package interfaces

import "github.com/norun9/Hybird/internal/api/interfaces/controllers"

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
	Func any
}

func GetMapRoute(messageController controllers.IMessageController) map[Path]Handler {
	return map[Path]Handler{
		{"/v1/messages", Get}:    {messageController.List},
		{"/v1/messages", Post}:   {messageController.Create},
		{"/v1/messages/ws", Get}: {messageController.Receive},
	}
}
