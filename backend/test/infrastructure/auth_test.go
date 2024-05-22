package infrastructure

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/auth"
	"MelkOnline/internal/infrastructure/signup"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
)

func Test_auth_inf(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	u := &model.User{
		Email:    "example@test.com",
		Password: "test",
		Name:     "test",
		Salt:     random.New().String(10),
		Type:     "test",
		Score:    0,
	}
	sr := signup.NewSignupRepository()
	_, err = sr.StoreUser(u.Email, u.Password, u.Name, u.Salt, u.Type, u.Score)
	assert.Nil(t, err, "Error should be nil")
	ar := auth.NewAuthRepository()
	user, err := ar.FindUserByEmail(u.Email)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, u.Email, user.Email, "Email should be equal")

}
