package signup

import (
	model "MelkOnline/internal/core"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SignupRepository struct {
	DBconn *gorm.DB
}

func NewSignupRepository() *SignupRepository {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &SignupRepository{DBconn: db}
}

func (sr *SignupRepository) StoreUser(email string, password string, name string, salt string, userType string, score float32) error {
	user := model.User{
		Email:    email,
		Password: password,
		Salt:     salt,
		Name:     name,
		Type:     userType,
		Score:    score,
	}
	err := sr.DBconn.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
