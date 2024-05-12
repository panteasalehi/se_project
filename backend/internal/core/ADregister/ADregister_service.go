package ADregister

import (
	"MelkOnline/internal/infrastructure/ADregister"
)

type ADregisterService struct {
	sr ADregisterRepositoryContract
}

func NewADregisterService() *ADregisterService {
	return &ADregisterService{
		sr: ADregister.NewADregisterRepository(),
	}
}

func (ss *ADregisterService) ADregister(title string, category string, price int, area float32, numberOfRooms int,
	yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int) error {
	err := ss.sr.StoreAD(title, category, price, area, numberOfRooms, yearOfConstruction, floor, description, elevator, store, parking, OwnerID)
	if err != nil {
		return err
	}
	return nil
}
