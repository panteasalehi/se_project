package chat

import (
	model "MelkOnline/internal/core"
)

type ChatContract interface {
	SendMessage(model.Message) (int, error)
	GetMessagesByChatID(chatID int) ([]model.Message, error)
}

type ChatRepositoryContract interface {
	StoreMessage(model.Message) (int, error)
	GetMessagesByChatID(chatID int) ([]model.Message, error)
}
