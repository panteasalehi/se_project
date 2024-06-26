package auth

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/auth"
	"crypto"
	"encoding/base64"
	"errors"

	"github.com/labstack/gommon/random"
)

type AuthService struct {
	arc AuthRepositoryContract
	sc  SessionContract
	rnd *random.Random
}

func NewAuthService() *AuthService {
	return &AuthService{
		arc: auth.NewAuthRepository(),
		sc:  auth.NewSessionRepository(),
		rnd: random.New(),
	}
}

func (as *AuthService) Login(email string, password string) (model.Session, error) {
	user, err := as.arc.FindUserByEmail(email)
	if err != nil {
		return model.Session{}, err
	}

	if !checkPassword(password, user.Salt, user.Password) {
		return model.Session{}, errors.New("invalid password")
	}

	token := as.GenerateToken()
	err = as.sc.StoreSession(token, user.ID)
	if err != nil {
		return model.Session{}, err
	}

	return model.Session{Token: token}, nil
}

func (as *AuthService) GenerateToken() string {
	return as.rnd.String(64)
}

func checkPassword(password string, salt string, passwordHash string) bool {
	return passwordHash == HashPassword(password, salt)
}

func HashPassword(password string, salt string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(password + salt))
	b64hash := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return b64hash
}
