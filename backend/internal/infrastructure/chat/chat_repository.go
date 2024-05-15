package chat

import (
	model "MelkOnline/internal/core"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ChatRepsitory struct {
	DBConn *gorm.DB
}

func NewChatRepository() *ChatRepsitory {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &ChatRepsitory{
		DBConn: db,
	}
}

func (cr *ChatRepsitory) StoreMessage(message model.Message) error {
	err := cr.DBConn.Create(&message).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *ChatRepsitory) GetMessagesByChatID(chatID int) ([]model.Message, error) {
	var messages []model.Message
	err := cr.DBConn.Where("chat_id = ?", chatID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
