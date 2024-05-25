package getpost

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"

	"gorm.io/gorm"
)

type GetPostRepository struct {
	DBconn *gorm.DB
}

func NewGetPostRepository() *GetPostRepository {
	return &GetPostRepository{DBconn: infrastructure.GetDB()}
}

func (gpr *GetPostRepository) GetPost(id int) (*core.AD, error) {
	var ad core.AD
	err := gpr.DBconn.Where("id = ?", id).First(&ad).Error
	if err != nil {
		return nil, err
	}
	return &ad, nil
}
