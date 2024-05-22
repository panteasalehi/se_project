package infrastructure

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/ADregister"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_ADreg_inf(t *testing.T) {
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
		YearOfConstruction: 1400,
		Floor:              "test",
		Description:        "test",
		Elevator:           true,
		Store:              true,
		Parking:            true,
		UserID:             1,
		Lt:                 1.0,
		Long:               1.0,
	}
	ADreg := ADregister.NewADregisterRepository()
	ID, err := ADreg.StoreAD(AD.Title, AD.Category, AD.Price, AD.Area, AD.NumberOfRooms, AD.YearOfConstruction, AD.Floor, AD.Description, AD.Elevator, AD.Store, AD.Parking, AD.UserID, AD.Lt, AD.Long)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}
