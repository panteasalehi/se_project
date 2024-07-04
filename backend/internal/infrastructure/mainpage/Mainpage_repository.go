package mainpage

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"

	"gorm.io/gorm"
)

type MainpageRepository struct {
	DBConn *gorm.DB
}

func NewMainpageRepository() *MainpageRepository {
	return &MainpageRepository{DBConn: infrastructure.GetDB()}
}

func (ar *MainpageRepository) FindADs() ([]model.AD, error) {
	ads := []model.AD{}
	err := ar.DBConn.Exec("SELECT * FROM ads").Scan(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}
