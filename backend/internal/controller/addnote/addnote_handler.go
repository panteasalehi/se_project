package addnote

import (
	"MelkOnline/internal/core"
	"MelkOnline/internal/core/addnote"
	"MelkOnline/internal/core/getpost"
	"MelkOnline/internal/infrastructure/auth"

	model "MelkOnline/internal/controller"

	"github.com/labstack/echo/v4"
)

type NoteHandler struct {
	ns addnote.NoteContract
	gc getpost.GetPostContract
	sr *auth.SessionRepository
}

func NewNoteHandler() *NoteHandler {
	return &NoteHandler{
		ns: addnote.NewNoteService(),
		gc: getpost.NewGetPostService(),
		sr: auth.NewSessionRepository(),
	}
}

func (nh *NoteHandler) StoreNote(c echo.Context) error {
	token, err := c.Request().Cookie("session")
	if err != nil {
		return c.JSON(400, "session not found")
	}
	userID, err := nh.sr.ValidateSession(token.Value)

	if err != nil {
		return c.JSON(400, "error in validating session")
	}
	var note core.UserNote
	if err := c.Bind(&note); err != nil {
		return c.JSON(400, "error in binding note")
	}
	note.UserID = userID
	_, err = nh.ns.StoreNote(note)
	if err != nil {
		return c.JSON(400, "error in storing note")
	}
	return c.JSON(200, "note sent")
}

func (nh *NoteHandler) GetNotes(c echo.Context) error {
	nreq := &model.NoteGetNotesRequest{}
	nres := &model.NoteGetNotesResponse{}
	session, err := c.Request().Cookie("session")
	if err != nil {
		nres.Message = "Session not found"
		return c.JSON(400, nres)
	}
	notes, err := nh.ns.GetNotsByUserID(nreq.UserID, session.Value)
	if err != nil {
		nres.Message = err.Error()
		return c.JSON(500, nres)
	}
	nres.Message = "Ads retrieved successfully"
	nres.Notes = notes
	return c.JSON(200, notes)
}
