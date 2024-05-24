package getpost

import (
	"MelkOnline/internal/core"
)

type GetPostContract interface {
	GetPost(int, string) (*core.AD, error)
}

type GetPostRepositoryContract interface {
	GetPost(int) (*core.AD, error)
}
