package chat

import (
	model "MelkOnline/internal/core"
)

type ChatContract interface {
	SendMessage(model.Message) error
	GetMessagesByChatID(chatID int) ([]model.Message, error)
}

type ChatRepositoryContract interface {
	StoreMessage(model.Message) error
	GetMessagesByChatID(chatID int) ([]model.Message, error)
}
