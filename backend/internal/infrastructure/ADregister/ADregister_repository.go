package ADregister

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "MelkOnline/internal/core"
)

type ADregisterRepository struct {
	DBconn *gorm.DB
}

func NewADregisterRepository() *ADregisterRepository {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &ADregisterRepository{DBconn: db}
}

func (sr *ADregisterRepository) StoreAD(Ad *model.AD) (int, error) {
	err := sr.DBconn.Create(&Ad).Error
	if err != nil {
		return 0, err
	}
	return Ad.ID, nil
}
