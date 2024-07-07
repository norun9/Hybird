package interfaces

import "github.com/norun9/Hybird/internal/api/usecase"

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

func GetMapRoute(messageInteractor usecase.IMessageInputBoundary) map[Path]Handler {
	return map[Path]Handler{
		{"/v1/messages", Get}: {messageInteractor.List},
	}
}
