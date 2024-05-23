package ADregister

import "MelkOnline/internal/core"

type ADregisterContract interface {
	ADregister(token string, title string, category string, price string, area string, numberOfRooms string,
		yearOfConstruction string, floor string, description string, elevator string, store string, parking string, OwnerID string, Lt float64,
		Long float64, AvatarURL string) (int, error)
}

type ADregisterRepositoryContract interface {
	StoreAD(Ad *core.AD) (int, error)
}
