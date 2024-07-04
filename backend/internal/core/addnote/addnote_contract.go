package addnote

import (
	model "MelkOnline/internal/core"
)

type NoteContract interface {
	StoreNote(note model.UserNote) (int, error)
	GetNotsByUserID(UserID int, token string) ([]model.UserNote, error)
}

type NoteRepositoryContract interface {
	StoreNote(note model.UserNote) (int, error)
	GetNotesByUserID(UserID int) ([]model.UserNote, error)
}
