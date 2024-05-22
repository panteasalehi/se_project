package core

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/auth"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_auth_valid_core(t *testing.T) {
	err := godotenv.Load("/home/ssaeidifarzad/ssfdata/ssaeidifarzad/Classes/S8/SE/Project/SE_project/backend/.env")
	if err != nil {
		panic(err)
	}
	user := &model.User{
		Email:    "galskjd",
		Password: "abcd",
		Name:     "abcd",
	}
	as := auth.NewAuthService()
	session, err := as.Login(user.Email, user.Password)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, "", session.Token, "Token should not be empty")
}

func Test_auth_invalid_core(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	user := &model.User{
		Email:    "galskjd",
		Password: "wrong",
		Name:     "abcd",
	}
	as := auth.NewAuthService()
	session, err := as.Login(user.Email, "wrong password")
	assert.NotNil(t, err, "Error should not be nil")
	assert.Equal(t, "", session.Token, "Token should be empty")
}
