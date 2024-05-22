package infrastructure

import (
	"MelkOnline/internal/infrastructure/mainpage"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_mainpage_inf(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	mr := mainpage.NewMainpageRepository()
	ADs, err := mr.FindADs()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(ADs))
}
