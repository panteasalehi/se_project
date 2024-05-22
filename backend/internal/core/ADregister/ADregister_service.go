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

func (ss *ADregisterService) ADregister(token string, title string, category string, price int, area float32, numberOfRooms int,
	yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int, Lt float32,
	Long float32, AvatarURL string) (int, error) {
	uid, err := ss.sr.ValidateSession(token)
	if err != nil || uid != OwnerID {
		return 0, err
	}
	ID, err := ss.arr.StoreAD(title, category, price, area, numberOfRooms, yearOfConstruction, floor, description, elevator, store,
		parking, OwnerID, Lt, Long, AvatarURL)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
