package input

import "github.com/norun9/Hybird/internal/api/domain/model"

type MessageList struct {
	Paging model.Paging `json:"paging" validate:"dive"`
}

// MessageInput call from Controller layer
type MessageInput struct {
	Content string `json:"content"`
}
