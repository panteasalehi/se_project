package chat

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"

	"gorm.io/gorm"
)

type ChatRepsitory struct {
	DBConn *gorm.DB
}

func NewChatRepository() *ChatRepsitory {
	return &ChatRepsitory{DBConn: infrastructure.GetDB()}
}

func (cr *ChatRepsitory) StoreMessage(message model.Message) (int, error) {
	err := cr.DBConn.Create(&message).Error
	if err != nil {
		return 0, err
	}
	return message.ID, nil
}

func (cr *ChatRepsitory) GetMessagesByChatID(chatID int) ([]model.Message, error) {
	var messages []model.Message
	err := cr.DBConn.Where("chat_id = ?", chatID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
