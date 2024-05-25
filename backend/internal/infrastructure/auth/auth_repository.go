package auth

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DBConn *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{DBConn: infrastructure.GetDB()}
}

func (ar *AuthRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := ar.DBConn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
