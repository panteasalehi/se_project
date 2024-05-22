package core

import (
	"testing"

	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/ADregister"
	"MelkOnline/internal/infrastructure/auth"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_ADregister_core(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	AD := &model.AD{
		Title:              "test",
		Category:           "test",
		Price:              100,
		Area:               100,
		NumberOfRooms:      1,
		YearOfConstruction: 1399,
		Floor:              "test",
		Description:        "test",
		Elevator:           true,
		Store:              true,
		Parking:            true,
		UserID:             1,
		Lt:                 1,
		Long:               1,
		AvatarURL:          "test",
	}
	as := ADregister.NewADregisterService()
	sr := auth.NewSessionRepository()
	err = sr.StoreSession("1", 1)
	if err != nil {
		panic(err)
	}
	ID, err := as.ADregister("1", AD.Title, AD.Category, AD.Price, AD.Area, AD.NumberOfRooms, AD.YearOfConstruction, AD.Floor, AD.Description, AD.Elevator, AD.Store, AD.Parking, AD.UserID, AD.Lt, AD.Long, AD.AvatarURL)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, ID)
}
