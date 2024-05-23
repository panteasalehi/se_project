package ADregister

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/core/auth"
	"MelkOnline/internal/infrastructure/ADregister"
	authrep "MelkOnline/internal/infrastructure/auth"
	"strconv"
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

func (ss *ADregisterService) ADregister(token string, title string, category string, price string, area string, numberOfRooms string,
	yearOfConstruction string, floor string, description string, elevator string, store string, parking string, OwnerID string, Lt string,
	Long string, AvatarURL string) (int, error) {
	uid, err := ss.sr.ValidateSession(token)
	if err != nil {
		return 0, err
	}
	AD, err := convertToAD(title, category, price, area, numberOfRooms, yearOfConstruction, floor, description, elevator, store, parking, Lt, Long, AvatarURL)
	if err != nil {
		return 0, err
	}
	AD.UserID = uid
	ID, err := ss.arr.StoreAD(AD)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func convertToAD(title string, category string, price string, area string, numberOfRooms string,
	yearOfConstruction string, floor string, description string, elevator string, store string, parking string, Lt string,
	Long string, AvatarURL string) (*core.AD, error) {
	AD := &core.AD{}
	var err error
	AD.Title = title
	AD.Category = category
	AD.Price, err = strconv.Atoi(price)
	if err != nil {
		return nil, err
	}
	AD.Area, err = strconv.ParseFloat(area, 64)
	if err != nil {
		return nil, err
	}
	AD.NumberOfRooms, err = strconv.Atoi(numberOfRooms)
	if err != nil {
		return nil, err
	}
	AD.YearOfConstruction, err = strconv.Atoi(yearOfConstruction)
	if err != nil {
		return nil, err
	}
	AD.Floor = floor
	AD.Description = description
	AD.Elevator, err = strconv.ParseBool(elevator)
	if err != nil {
		return nil, err
	}
	AD.Store, err = strconv.ParseBool(store)
	if err != nil {
		return nil, err
	}
	AD.Parking, err = strconv.ParseBool(parking)
	if err != nil {
		return nil, err
	}
	AD.Lt, err = strconv.ParseFloat(Lt, 64)
	if err != nil {
		return nil, err
	}
	AD.Long, err = strconv.ParseFloat(Long, 64)
	if err != nil {
		return nil, err
	}
	AD.AvatarURL = AvatarURL
	return AD, nil
}
