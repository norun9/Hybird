package controller

import "github.com/norun9/Hybird/internal/api/usecase"

type MessageController struct {
	interactor usecase.IMessageInputBoundary
}

func NewMessageController(interactor usecase.IMessageInputBoundary) *MessageController {
	return &MessageController{interactor}
}
