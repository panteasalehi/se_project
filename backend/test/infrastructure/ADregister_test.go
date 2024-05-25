package infrastructure

import (
	"MelkOnline/internal"
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/ADregister"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ADreg_inf(t *testing.T) {
	es := internal.NewEchoServer()
	go es.Start()
	time.Sleep(1 * time.Second)
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
		AvatarURL:          "test",
	}
	ADreg := ADregister.NewADregisterRepository()
	ID, err := ADreg.StoreAD(AD)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}
