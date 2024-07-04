package addnote

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"

	"gorm.io/gorm"
)

type NoteRepsitory struct {
	DBConn *gorm.DB
}

func NewNoteRepository() *NoteRepsitory {
	return &NoteRepsitory{DBConn: infrastructure.GetDB()}
}

func (nr *NoteRepsitory) StoreNote(note model.UserNote) (int, error) {
	err := nr.DBConn.Create(&note).Error
	if err != nil {
		return 0, err
	}
	return note.ID, nil
}

func (nr *NoteRepsitory) GetNotesByUserID(UserID int) ([]model.UserNote, error) {
	var notes []model.UserNote
	err := nr.DBConn.Where("user_id = ?", UserID).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}
