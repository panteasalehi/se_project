package chat

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/chat"
)

type ChatService struct {
	cr ChatRepositoryContract
}

func NewChatService() *ChatService {
	return &ChatService{
		cr: chat.NewChatRepository(),
	}
}

func (cs *ChatService) SendMessage(message model.Message) error {
	err := cs.cr.StoreMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChatService) GetMessagesByChatID(chatID int) ([]model.Message, error) {
	messages, err := cs.cr.GetMessagesByChatID(chatID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
