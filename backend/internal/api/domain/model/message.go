package model

import "time"

type Message struct {
	ID        int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
