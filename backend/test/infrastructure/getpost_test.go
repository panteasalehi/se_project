package infrastructure

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/ADregister"
	"MelkOnline/internal/infrastructure/getpost"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Getpost_inf(t *testing.T) {
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
		Lt:                 1.0,
		Long:               1.0,
		AvatarURL:          "test",
	}
	adr := ADregister.NewADregisterRepository()
	ID, err := adr.StoreAD(AD)
	if err != nil {
		panic(err)
	}
	gpr := getpost.NewGetPostRepository()
	ad, err := gpr.GetPost(ID)
	assert.Nil(t, err)
	assert.NotNil(t, ad)
	assert.Equal(t, "test", ad.Title)
	assert.Equal(t, "test", ad.Category)
	assert.Equal(t, 100, ad.Price)
	assert.Equal(t, 100.0, ad.Area)
	assert.Equal(t, 1, ad.NumberOfRooms)
	assert.Equal(t, 1399, ad.YearOfConstruction)
	assert.Equal(t, "test", ad.Floor)
	assert.Equal(t, "test", ad.Description)
	assert.Equal(t, true, ad.Elevator)
	assert.Equal(t, true, ad.Store)
	assert.Equal(t, true, ad.Parking)
	assert.Equal(t, 1, ad.UserID)
	assert.Equal(t, 1.0, ad.Lt)
	assert.Equal(t, 1.0, ad.Long)
	assert.Equal(t, "test", ad.AvatarURL)
}
