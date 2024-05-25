package getpost

import (
	"MelkOnline/internal/core"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GetPostRepository struct {
	DBconn *gorm.DB
}

func NewGetPostRepository() *GetPostRepository {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &GetPostRepository{DBconn: db}
}

func (gpr *GetPostRepository) GetPost(id int) (*core.AD, error) {
	var ad core.AD
	err := gpr.DBconn.Where("id = ?", id).First(&ad).Error
	if err != nil {
		return nil, err
	}
	return &ad, nil
}
