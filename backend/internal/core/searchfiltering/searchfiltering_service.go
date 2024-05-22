package searchfiltering

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/searchfiltering"
)

type SearchFilteringService struct {
	cr SeatchFilteringRepositoryContract
}

func NewSearchFilteringService() *SearchFilteringService {
	return &SearchFilteringService{
		cr: searchfiltering.NewSearchFilteringRepository(),
	}
}

func (cs *SearchFilteringService) FindCategoryFilteredADs(Category string) ([]model.AD, error) {
	ads, err := cs.cr.FindCategoryFilteredADs(Category)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (cs *SearchFilteringService) FindPriceFilteredADs(MinPrice int, MaxPrice int) ([]model.AD, error) {
	ads, err := cs.cr.FindPriceFilteredADs(MinPrice, MaxPrice)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (cs *SearchFilteringService) FindAreaFilteredADs(Minarea int, Maxarea int) ([]model.AD, error) {
	ads, err := cs.cr.FindPriceFilteredADs(Minarea, Maxarea)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (cs *SearchFilteringService) FindRoomFilteredADs(NumberOfRooms int) ([]model.AD, error) {
	ads, err := cs.cr.FindRoomFilteredADs(NumberOfRooms)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (cs *SearchFilteringService) FindPSEFilteredADs(Parking bool, Store bool, Elevator bool) ([]model.AD, error) {
	ads, err := cs.cr.FindPSEFilteredADs(Parking, Store, Elevator)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (cs *SearchFilteringService) FindFloorFilteredADs(MinFloor int, MaxFloor int) ([]model.AD, error) {
	ads, err := cs.cr.FindFloorFilteredADs(MinFloor, MaxFloor)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (cs *SearchFilteringService) FindYearFilteredADs(MaxAge int) ([]model.AD, error) {
	ads, err := cs.cr.FindYearFilteredADs(MaxAge)
	if err != nil {
		return nil, err
	}
	return ads, nil
}
