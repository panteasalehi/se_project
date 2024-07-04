package user

import "MelkOnline/internal/core"

type UserServiceContract interface {
	GetUserBySession(session string) (*core.User, error)
	GetUserAds(session string) ([]core.AD, error)
}

type UserRepositryContract interface {
	GetUserBySession(session string) (*core.User, error)
	GetUserAds(session string) ([]core.AD, error)
}
