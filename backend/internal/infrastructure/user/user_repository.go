package user

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"
	"MelkOnline/internal/infrastructure/auth"

	"gorm.io/gorm"
)

type UserRepositry struct {
	dbconn *gorm.DB
	sr     *auth.SessionRepository
}

func NewUserRepositry() *UserRepositry {
	return &UserRepositry{
		dbconn: infrastructure.GetDB(),
		sr:     auth.NewSessionRepository(),
	}
}

func (ur *UserRepositry) GetUserBySession(session string) (*core.User, error) {
	ID, err := ur.sr.ValidateSession(session)
	if err != nil {
		return nil, err
	}
	user := &core.User{}
	ur.dbconn.First(user, ID)
	return user, nil
}

func (ur *UserRepositry) GetUserAds(session string) ([]core.AD, error) {
	ID, err := ur.sr.ValidateSession(session)
	if err != nil {
		return nil, err
	}
	ads := []core.AD{}
	ur.dbconn.Find(&ads, "user_id = ?", ID)
	return ads, nil
}
