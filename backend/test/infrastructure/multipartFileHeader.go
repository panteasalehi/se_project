package infrastructure

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func CreateMultipartFileHeader(filePath string) *multipart.FileHeader {
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
