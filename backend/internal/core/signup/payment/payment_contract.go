package payment

type PaymentServiceContract interface {
	Pay(userID int, amount int, status string) (int, error)
	GetLastID() (int, error)
}

type PaymentRepositoryContract interface {
	StorePayment(userID int, amount int, status string) (int, error)
	GetLastID() (int, error)
}
