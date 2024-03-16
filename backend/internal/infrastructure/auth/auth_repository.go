package auth

import (
	model "MelkOnline/internal/core"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DBConn *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	db, err := gorm.Open(mysql.Open("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
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
