package Mainpage

import (
	model "MelkOnline/internal/core"
	mainpage "MelkOnline/internal/infrastructure/mainpage"
)

type MainpageService struct {
	sr MainpageRepositoryContract
}

func NewMainpageService() *MainpageService {
	return &MainpageService{
		sr: mainpage.NewMainpageRepository(),
	}
}

func (ss *MainpageService) Mainpage() ([]model.AD, error) {
	ADs, err := ss.sr.FindADs()
	if err != nil {
		return nil, err
	}
	return ADs, nil
}
