package infrastructure

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/ADregister"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ADreg_inf(t *testing.T) {
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
	image := CreateMultipartFileHeader("/home/runner/work/se_project/se_project/backend/test/infrastructure/test.jpg")
	ID, err := ADreg.StoreAD(AD, image)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}
