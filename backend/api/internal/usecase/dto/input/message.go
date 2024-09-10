package input

import "github.com/norun9/HyBird/backend/api/internal/domain/model"

type MessageList struct {
	Paging model.Paging `json:"paging"`
}

type MessageInput struct {
	Content string `json:"content"`
}
