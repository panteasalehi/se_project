package ADregister

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "MelkOnline/internal/core"
)

type ADregisterRepository struct {
	DBconn *gorm.DB
}

func NewADregisterRepository() *ADregisterRepository {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &ADregisterRepository{DBconn: db}
}

func (sr *ADregisterRepository) StoreAD(title string, category string, price int, area float32, numberOfRooms int,
	yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int) (int, error) {
	Ad := model.AD{
		Title:              title,
		Category:           category,
		Price:              price,
		Area:               area,
		NumberOfRooms:      numberOfRooms,
		YearOfConstruction: yearOfConstruction,
		Floor:              floor,
		Description:        description,
		Elevator:           elevator,
		Store:              store,
		Parking:            parking,
		UserID:             OwnerID,
		//location
	}
	err := sr.DBconn.Create(&Ad).Error
	if err != nil {
		return 0, err
	}
	return Ad.ID, nil
}
