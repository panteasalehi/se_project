package ADregister

type ADregisterContract interface {
	ADregister(title string, category string, price int, area float32, numberOfRooms int,
		yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int) error
}

type ADregisterRepositoryContract interface {
	StoreAD(title string, category string, price int, area float32, numberOfRooms int,
		yearOfConstruction int, floor string, description string, elevator bool, store bool, parking bool, OwnerID int) error
}
