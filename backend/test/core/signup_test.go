package core

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/signup"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_signup_service(t *testing.T) {
	err := godotenv.Load("/home/ssaeidifarzad/ssfdata/ssaeidifarzad/Classes/S8/SE/Project/SE_project/backend/.env")
	if err != nil {
		panic(err)
	}
	user := &model.User{
		Email:    "galskjd",
		Password: "abcd",
		Name:     "abcd",
	}
	sr := signup.NewSignupService()
	ID, err := sr.Signup(user.Email, user.Password, user.Name)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}
