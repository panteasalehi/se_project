package payment

import "MelkOnline/internal/infrastructure/signup/payment"

type PaymentService struct {
	pr PaymentRepositoryContract
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		pr: payment.NewPaymentRepository(),
	}
}

func (ps *PaymentService) Pay(userID int, amount int, status string) (int, error) {
	return ps.pr.StorePayment(userID, amount, status)
}

func (ps *PaymentService) GetLastID() (int, error) {
	return ps.pr.GetLastID()
}
