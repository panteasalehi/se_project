package infrastructure

import (
	"MelkOnline/internal/infrastructure/mainpage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mainpage_inf(t *testing.T) {
	mr := mainpage.NewMainpageRepository()
	ADs, err := mr.FindADs()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(ADs))
}
