package chat

import (
	model "MelkOnline/internal/core"
)

type ChatContract interface {
	SendMessage(model.Message) (int, error)
	GetMessagesByChatID(chatID int) ([]model.Message, error)
	CreateChat(user1, user2, ad int) (int, error)
	ChatExists(adID, userID int) (bool, error)
}

type ChatRepositoryContract interface {
	StoreChat(chat model.Chat) (int, error)
	StoreMessage(model.Message) (int, error)
	GetMessagesByChatID(chatID int) ([]model.Message, error)
	ChatExists(adID, userID int) (bool, error)
}
