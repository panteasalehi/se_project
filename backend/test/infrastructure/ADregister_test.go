package infrastructure

import (
	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure/ADregister"
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ADreg_inf(t *testing.T) {
	AD := &model.AD{
		Title:              "test",
		Category:           "test",
		Price:              100,
		Area:               100,
		NumberOfRooms:      1,
		YearOfConstruction: 1400,
		Floor:              "test",
		Description:        "test",
		Elevator:           true,
		Store:              true,
		Parking:            true,
		UserID:             1,
		Lt:                 1.0,
		Long:               1.0,
	}
	ADreg := ADregister.NewADregisterRepository()
	err := os.Mkdir("images", 0755)
	if err != nil {
		panic(err)
	}
	image := createMultipartFileHeader("/home/runner/work/se_project/se_project/backend/test/infrastructure/test.jpg")
	ID, err := ADreg.StoreAD(AD, image)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, ID, "ID should not be 0")
}

func createMultipartFileHeader(filePath string) *multipart.FileHeader {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var buff bytes.Buffer
	buffWriter := io.Writer(&buff)
	formWriter := multipart.NewWriter(buffWriter)
	formPart, err := formWriter.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(formPart, file); err != nil {
		panic(err)
	}
	formWriter.Close()
	buffReader := bytes.NewReader(buff.Bytes())
	formReader := multipart.NewReader(buffReader, formWriter.Boundary())
	multipartForm, err := formReader.ReadForm(1 << 20)
	if err != nil {
		panic(err)
	}
	files, exists := multipartForm.File["file"]
	if !exists || len(files) == 0 {
		log.Fatal("multipart file not exists")
		return nil
	}

	return files[0]
}
