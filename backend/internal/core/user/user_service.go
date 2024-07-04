package user

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/user"
)

type UserService struct {
	ur UserRepositryContract
}

func NewUserService() *UserService {
	return &UserService{
		ur: user.NewUserRepositry(),
	}
}

func (us *UserService) GetUserBySession(session string) (*core.User, error) {
	return us.ur.GetUserBySession(session)
}

func (us *UserService) GetUserAds(session string) ([]core.AD, error) {
	return us.ur.GetUserAds(session)
}
