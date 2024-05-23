package infrastructure

import (
	model "MelkOnline/internal/core"
	"testing"

	signup "MelkOnline/internal/infrastructure/signup"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
)

func Test_signup_inf(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	u := &model.User{
		Email:    "test",
		Password: "test",
		Name:     "test",
		Salt:     random.New().String(10),
		Type:     "test",
		Score:    0,
	}
	sr := signup.NewSignupRepository()
	ID, err := sr.StoreUser(u.Email, u.Password, u.Name, u.Salt, u.Type, u.Score)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}
