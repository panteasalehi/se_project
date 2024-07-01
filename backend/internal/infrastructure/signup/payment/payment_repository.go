package payment

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{db: infrastructure.GetDB()}
}

func (pr *PaymentRepository) StorePayment(userID int, amount int, status string) (int, error) {
	payment := &core.Payment{
		UserID: userID,
		Amount: amount,
		Status: status,
	}
	pr.db.Create(payment)
	return payment.ID, nil
}

func (pr *PaymentRepository) GetLastID() (int, error) {
	var payment core.Payment
	pr.db.Last(&payment)
	return payment.ID, nil
}
