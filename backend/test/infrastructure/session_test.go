package infrastructure

import (
	"MelkOnline/internal/infrastructure/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_session_token_not_valid(t *testing.T) {
	sr := auth.NewSessionRepository()
	token := "invalid"
	valid, err := sr.ValidateSession(token)
	assert.NotNil(t, err, "Error should not be nil")
	assert.False(t, valid, "Session should not be valid")
}

func Test_session_set_token(t *testing.T) {
	sr := auth.NewSessionRepository()
	token := "valid"
	userID := 1
	err := sr.StoreSession(token, userID)
	assert.Nil(t, err, "Error should be nil")
}

func Test_session_token_valid(t *testing.T) {
	sr := auth.NewSessionRepository()
	token := "valid"
	valid, err := sr.ValidateSession(token)
	assert.Nil(t, err, "Error should be nil")
	assert.True(t, valid, "Session should be valid")
}
