package core

import (
	"testing"

	model "MelkOnline/internal/controller"
	"MelkOnline/internal/core/ADregister"
	"MelkOnline/internal/infrastructure/auth"
	testinf "MelkOnline/test/infrastructure"

	"github.com/stretchr/testify/assert"
)

func Test_ADregister_core(t *testing.T) {
	AD := &model.ADregisterRequest{
		Title:              "test",
		Category:           "test",
		Price:              "100",
		Area:               "100",
		NumberOfRooms:      "1",
		YearOfConstruction: "1399",
		Floor:              "test",
		Description:        "test",
		Elevator:           "true",
		Store:              "true",
		Parking:            "true",
		OwnerID:            "1",
		Lt:                 1.0,
		Long:               1.0,
		Image:              "test.jpg",
	}
	as := ADregister.NewADregisterService()
	sr := auth.NewSessionRepository()
	err := sr.StoreSession("1", 1)
	if err != nil {
		panic(err)
	}
	image := testinf.CreateMultipartFileHeader("/home/runner/work/se_project/se_project/backend/test/infrastructure/test.jpg")
	ID, err := as.ADregister("1", AD.Title, AD.Category, AD.Price, AD.Area, AD.NumberOfRooms, AD.YearOfConstruction, AD.Floor, AD.Description, AD.Elevator, AD.Store, AD.Parking, AD.OwnerID, AD.Lt, AD.Long, image)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, ID)
}
