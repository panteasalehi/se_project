package ADregister

type ADregisterContract interface {
	ADregister(token string, title string, category string, price int, area float32, numberOfRooms int,
		yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int, Lt float32,
		Long float32) (int, error)
}

type ADregisterRepositoryContract interface {
	StoreAD(title string, category string, price int, area float32, numberOfRooms int,
		yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int, Lt float32,
		Long float32) (int, error)
}
