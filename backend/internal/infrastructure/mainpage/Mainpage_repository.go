package mainpage

import (
	model "MelkOnline/internal/core"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MainpageRepository struct {
	DBConn *gorm.DB
}

func NewMainpageRepository() *MainpageRepository {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &MainpageRepository{DBConn: db}
}

func (ar *MainpageRepository) FindADs() ([]model.AD, error) {
	ads := []model.AD{}
	err := ar.DBConn.Model(&model.AD{}).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}
