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

func (cs *ChatService) CreateChat(user1, user2, ad int) (int, error) {
	chat := model.Chat{
		User1ID: user1,
		User2ID: user2,
		AdID:    ad,
	}
	ID, err := cs.cr.StoreChat(chat)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func (cs *ChatService) SendMessage(message model.Message) (int, error) {
	ID, err := cs.cr.StoreMessage(message)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func (cs *ChatService) GetMessagesByChatID(chatID int) ([]model.Message, error) {
	messages, err := cs.cr.GetMessagesByChatID(chatID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (cs *ChatService) ChatExists(adID, userID int) (int, error) {
	return cs.cr.ChatExists(adID, userID)
}

func (cs *ChatService) GetChatInfo(chatID int) (model.Chat, error) {
	chat, err := cs.cr.GetChatInfo(chatID)
	if err != nil {
		return model.Chat{}, err
	}
	return chat, nil
}
