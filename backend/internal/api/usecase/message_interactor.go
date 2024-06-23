package usecase

import (
	"github.com/norun9/Hybird/internal/api/usecase/dto/input"
	"github.com/norun9/Hybird/internal/api/usecase/dto/output"
	"gorm.io/gorm"
)

type MessageInputBoundary interface {
	Create(input input.MessageInput) error
	List() ([]output.MessageOutput, error)
}

// NOTE:OutputBoundary interface definition is omitted to prevent code bloat.

type MessageInteractor struct {
	dbClient gorm.DB
}

func NewMessageInteractor(dbClient gorm.DB) MessageInputBoundary {
	return &MessageInteractor{dbClient}
}

func (i *MessageInteractor) Create(input input.MessageInput) error {
	return nil
}

func (i *MessageInteractor) List() ([]output.MessageOutput, error) {
	return nil, nil
}
