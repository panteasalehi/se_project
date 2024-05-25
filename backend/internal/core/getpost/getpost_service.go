package getpost

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/core/auth"
	authrep "MelkOnline/internal/infrastructure/auth"
	"MelkOnline/internal/infrastructure/getpost"
)

type GetPostService struct {
	repo GetPostRepositoryContract
	sr   auth.SessionContract
}

func NewGetPostService() *GetPostService {
	return &GetPostService{
		repo: getpost.NewGetPostRepository(),
		sr:   authrep.NewSessionRepository(),
	}
}

func (gps *GetPostService) GetPost(id int, token string) (*core.AD, error) {
	_, err := gps.sr.ValidateSession(token)
	if err != nil {
		return nil, err
	}
	ad, err := gps.repo.GetPost(id)
	if err != nil {
		return nil, err
	}
	return ad, nil
}
