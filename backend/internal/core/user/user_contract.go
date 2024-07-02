package user

import "MelkOnline/internal/core"

type UserServiceContract interface {
	GetUserBySession(session string) (*core.User, error)
}

type UserRepositryContract interface {
	GetUserBySession(session string) (*core.User, error)
}
