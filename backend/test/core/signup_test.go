package core

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/signup"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_signup_service(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	user := &model.User{
		Email:    "galskjd",
		Password: "abcd",
		Name:     "abcd",
	}
	sr := signup.NewSignupService()
	ID,token, err := sr.Signup(user.Email, user.Password, user.Name)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, "", token, "Token should not be empty")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}
