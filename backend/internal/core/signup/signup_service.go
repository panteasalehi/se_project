package signup

import (
	"MelkOnline/internal/core/auth"
	authrep "MelkOnline/internal/infrastructure/auth"
	"MelkOnline/internal/infrastructure/signup"
	"errors"

	"github.com/labstack/gommon/random"
)

type SignupService struct {
	sr SignupRepositoryContract
	ar auth.AuthRepositoryContract
}

func NewSignupService() *SignupService {
	return &SignupService{
		sr: signup.NewSignupRepository(),
		ar: authrep.NewAuthRepository(),
	}
}

func (ss *SignupService) Signup(email string, password string, name string) (int, error) {
	_, err := ss.ar.FindUserByEmail(email)
	if err == nil {
		return 0, errors.New("email already exists")
	}
	salt := random.New().String(10)
	userType := "user"
	score := float32(0)
	ID, err := ss.sr.StoreUser(email, password, name, salt, userType, score)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
