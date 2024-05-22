package searchfiltering

import (
	model "MelkOnline/internal/core"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SearchFilteringRepository struct {
	DBConn *gorm.DB
}

func NewSearchFilteringRepository() *SearchFilteringRepository {
	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &SearchFilteringRepository{DBConn: db}
}

func (ar *SearchFilteringRepository) FindCategoryFilteredADs(Category string) ([]model.AD, error) {
	var ads []model.AD
	err := ar.DBConn.Where("Category = ?", Category).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (ar *SearchFilteringRepository) FindPriceFilteredADs(Minprice int, Maxprice int) ([]model.AD, error) {
	var ads []model.AD
	err := ar.DBConn.Where("Price BETWEEN ? AND ?", Minprice, Maxprice).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (ar *SearchFilteringRepository) FindRoomFilteredADs(NumberOfRooms int) ([]model.AD, error) {
	var ads []model.AD
	err := ar.DBConn.Where("NumberOfRooms = ?", NumberOfRooms).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (ar *SearchFilteringRepository) FindAreaFilteredADs(Minarea int, Maxarea int) ([]model.AD, error) {
	var ads []model.AD
	err := ar.DBConn.Where("Area BETWEEN ? AND ?", Minarea, Maxarea).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (ar *SearchFilteringRepository) FindPSEFilteredADs(Parking bool, Store bool, Elevator bool) ([]model.AD, error) {
	var ads []model.AD
	err := ar.DBConn.Where("Parking = ? AND Store = ? AND Elevator = ?", Parking, Store, Elevator).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (ar *SearchFilteringRepository) FindFloorFilteredADs(MinFloor int, MaxFloor int) ([]model.AD, error) {
	var ads []model.AD
	err := ar.DBConn.Where("Floor BETWEEN ? AND ?", MinFloor, MaxFloor).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (ar *SearchFilteringRepository) FindYearFilteredADs(MaxAge int) ([]model.AD, error) {
	var ads []model.AD
	currentTime := time.Now()
	currentYear := currentTime.Year()
	currentYear -= 621
	year := currentYear - MaxAge
	err := ar.DBConn.Where("YearOfConstruction >= ?", year).Find(&ads).Error
	if err != nil {
		return nil, err
	}
	return ads, nil
}
