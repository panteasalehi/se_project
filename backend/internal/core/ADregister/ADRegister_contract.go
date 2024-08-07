package ADregister

import (
	"MelkOnline/internal/core"
	"mime/multipart"
)

type ADregisterContract interface {
	ADregister(token string, title string, category string, price string, area string, numberOfRooms string,
		yearOfConstruction string, floor string, description string, elevator string, store string, parking string, OwnerID string, Lt float64,
		Long float64, image *multipart.FileHeader) (int, error)
}

type ADregisterRepositoryContract interface {
	StoreAD(Ad *core.AD, image *multipart.FileHeader) (int, error)
}
