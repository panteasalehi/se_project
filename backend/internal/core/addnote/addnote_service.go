package addnote

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/core/auth"
	"MelkOnline/internal/infrastructure/addnote"
	authrep "MelkOnline/internal/infrastructure/auth"
)

type NoteService struct {
	nr NoteRepositoryContract
	sr auth.SessionContract
}

func NewNoteService() *NoteService {
	return &NoteService{
		nr: addnote.NewNoteRepository(),
		sr: authrep.NewSessionRepository(),
	}
}

func (ns *NoteService) StoreNote(note model.UserNote) (int, error) {
	ID, err := ns.nr.StoreNote(note)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func (ns *NoteService) GetNotsByUserID(UserID int, token string) ([]model.UserNote, error) {
	_, err := ns.sr.ValidateSession(token)
	if err != nil {
		return nil, err
	}

	notes, err := ns.nr.GetNotesByUserID(UserID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
