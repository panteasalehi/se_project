package Mainpage

import (
	model "MelkOnline/internal/core"
)

type MainpageContract interface {
	Mainpage() ([]model.AD, error)
}

type MainpageRepositoryContract interface {
	FindADs() ([]model.AD, error)
}
