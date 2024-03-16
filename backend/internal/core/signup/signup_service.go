package signup

import (
	"MelkOnline/internal/infrastructure/signup"

	"github.com/labstack/gommon/random"
)

type SignupService struct {
	sr SignupRepositoryContract
}

func NewSignupService() *SignupService {
	return &SignupService{
		sr: signup.NewSignupRepository(),
	}
}

func (ss *SignupService) Signup(email string, password string, name string) error {
	salt := random.New().String(10)
	userType := "user"
	score := float32(0)
	err := ss.sr.StoreUser(email, password, name, salt, userType, score)
	if err != nil {
		return err
	}
	return nil
}
