package searchfiltering

import (
	model "MelkOnline/internal/core"
)

type SearchFilteringContract interface { ///change
	FindCategoryFilteredADs(Category string) ([]model.AD, error)
	FindPriceFilteredADs(Minprice int, Maxprice int) ([]model.AD, error)
	FindAreaFilteredADs(Minarea int, Maxarea int) ([]model.AD, error)
	FindRoomFilteredADs(NumberOfRooms int) ([]model.AD, error)
	FindPSEFilteredADs(Parking bool, Store bool, Elevator bool) ([]model.AD, error)
	FindFloorFilteredADs(MinFloor int, MaxFloor int) ([]model.AD, error)
	FindYearFilteredADs(MaxAge int) ([]model.AD, error)
}

type SeatchFilteringRepositoryContract interface {
	FindCategoryFilteredADs(Category string) ([]model.AD, error)
	FindPriceFilteredADs(Minprice int, Maxprice int) ([]model.AD, error)
	FindAreaFilteredADs(Minarea int, Maxarea int) ([]model.AD, error)
	FindRoomFilteredADs(NumberOfRooms int) ([]model.AD, error)
	FindPSEFilteredADs(Parking bool, Store bool, Elevator bool) ([]model.AD, error)
	FindFloorFilteredADs(MinFloor int, MaxFloor int) ([]model.AD, error)
	FindYearFilteredADs(MaxAge int) ([]model.AD, error)
}
