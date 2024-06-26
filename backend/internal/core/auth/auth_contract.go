package auth

import model "MelkOnline/internal/core"

type AuthContract interface {
	Login(email string, password string) (model.Session, error)
}

type AuthRepositoryContract interface {
	FindUserByEmail(email string) (*model.User, error)
}

type SessionContract interface {
	ValidateSession(token string) (int, error)
	StoreSession(token string, userID int) error
}
