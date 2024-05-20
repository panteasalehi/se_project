package Mainpage

import (
	model "MelkOnline/internal/core"
	authcore "MelkOnline/internal/core/auth"
	"MelkOnline/internal/infrastructure/auth"
	mainpage "MelkOnline/internal/infrastructure/mainpage"
)

type MainpageService struct {
	sr   MainpageRepositoryContract
	srep authcore.SessionContract
}

func NewMainpageService() *MainpageService {
	return &MainpageService{
		sr:   mainpage.NewMainpageRepository(),
		srep: auth.NewSessionRepository(),
	}
}

func (ss *MainpageService) Mainpage(token string) ([]model.AD, error) {
	_, err := ss.srep.ValidateSession(token)
	if err != nil {
		return nil, err
	}
	ADs, err := ss.sr.FindADs()
	if err != nil {
		return nil, err
	}
	return ADs, nil
}
