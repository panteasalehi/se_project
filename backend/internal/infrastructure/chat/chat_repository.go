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

func (cr *ChatRepsitory) StoreChat(chat model.Chat) (int, error) {
	err := cr.DBConn.Create(&chat).Error
	if err != nil {
		return 0, err
	}
	return chat.ID, nil
}

func (cr *ChatRepsitory) ChatExists(adID, userID int) (bool, error) {
	Ad := model.AD{}
	err := cr.DBConn.Where("id = ?", adID).First(&Ad).Error
	if err != nil {
		return false, err
	}
	user1ID := Ad.UserID
	var chat model.Chat
	err = cr.DBConn.Where("ad_id = ? AND (user1_id = ? AND user2_id = ?)", adID, user1ID, userID).First(&chat).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
