package core

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/signup"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_signup_service(t *testing.T) {
	user := &model.User{
		Email:    "test1@example.com",
		Password: "abcd",
		Name:     "abcd",
	}
	sr := signup.NewSignupService()
	ID, token, err := sr.Signup(user.Email, user.Password, user.Name)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, "", token, "Token should not be empty")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}
