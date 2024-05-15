package auth

import (
	model "MelkOnline/internal/core"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DBConn *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &AuthRepository{DBConn: db}
}

func (ar *AuthRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := ar.DBConn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
