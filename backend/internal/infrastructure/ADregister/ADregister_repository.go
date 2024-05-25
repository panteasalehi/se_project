package ADregister

import (
	"gorm.io/gorm"

	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"
)

type ADregisterRepository struct {
	DBconn *gorm.DB
}

func NewADregisterRepository() *ADregisterRepository {
	return &ADregisterRepository{DBconn: infrastructure.GetDB()}
}

func (sr *ADregisterRepository) StoreAD(Ad *model.AD) (int, error) {
	err := sr.DBconn.Create(&Ad).Error
	if err != nil {
		return 0, err
	}
	return Ad.ID, nil
}
