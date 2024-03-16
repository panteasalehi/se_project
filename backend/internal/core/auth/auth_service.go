package auth

import (
	model "MelkOnline/internal/core"
	"errors"
)

type AuthService struct {
	arc AuthRepositoryContract
	sc  SessionContract
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Login(email string, password string) (model.Session, error) {
	user, err := as.arc.FindUserByEmail(email)
	if err != nil {
		return model.Session{}, err
	}

	if user.Password != password {
		return model.Session{}, errors.New("invalid password")
	}

	token := as.generateToken()
	err = as.sc.StoreSession(token, user.ID)
	if err != nil {
		return model.Session{}, err
	}

	return model.Session{Token: token}, nil
}

func (as *AuthService) generateToken() string {
	return "token"
}
