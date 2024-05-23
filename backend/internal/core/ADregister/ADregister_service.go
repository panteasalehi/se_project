package ADregister

import (
	"MelkOnline/internal/core/auth"
	"MelkOnline/internal/infrastructure/ADregister"
	authrep "MelkOnline/internal/infrastructure/auth"
)

type ADregisterService struct {
	arr ADregisterRepositoryContract
	sr  auth.SessionContract
}

func NewADregisterService() *ADregisterService {
	return &ADregisterService{
		arr: ADregister.NewADregisterRepository(),
		sr:  authrep.NewSessionRepository(),
	}
}

func (ss *ADregisterService) ADregister(token string, title string, category string, price int, area float64, numberOfRooms int,
	yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int, Lt float64,
	Long float64, AvatarURL string) (int, error) {
	uid, err := ss.sr.ValidateSession(token)
	if err != nil {
		return 0, err
	}
	ID, err := ss.arr.StoreAD(title, category, price, area, numberOfRooms, yearOfConstruction, floor, description, elevator, store,
		parking, uid, Lt, Long, AvatarURL)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
