package signup

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/auth"
	"MelkOnline/internal/infrastructure"

	"gorm.io/gorm"
)

type SignupRepository struct {
	DBconn *gorm.DB
}

func NewSignupRepository() *SignupRepository {
	return &SignupRepository{DBconn: infrastructure.GetDB()}
}

func (sr *SignupRepository) StoreUser(email string, password string, name string, salt string, userType string, score float32) (int, error) {
	user := &model.User{
		Email:    email,
		Password: auth.HashPassword(password, salt),
		Salt:     salt,
		Name:     name,
		Type:     userType,
		Score:    score,
	}
	err := sr.DBconn.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
