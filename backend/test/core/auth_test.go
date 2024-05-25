package core

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/auth"
	"MelkOnline/internal/infrastructure/signup"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_auth_valid_core(t *testing.T) {
	user := &model.User{
		Email:    "test@example.com",
		Password: "abcd",
		Name:     "abcd",
	}
	sr := signup.NewSignupRepository()
	sr.StoreUser(user.Email, user.Password, user.Name, auth.NewAuthService().GenerateToken(), "user", 0)
	as := auth.NewAuthService()
	session, err := as.Login(user.Email, user.Password)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, "", session.Token, "Token should not be empty")
}

func Test_auth_invalid_core(t *testing.T) {
	user := &model.User{
		Email:    "test@example.com",
		Password: "wrong",
		Name:     "abcd",
	}
	as := auth.NewAuthService()
	session, err := as.Login(user.Email, "wrong password")
	assert.NotNil(t, err, "Error should not be nil")
	assert.Equal(t, "", session.Token, "Token should be empty")
}
