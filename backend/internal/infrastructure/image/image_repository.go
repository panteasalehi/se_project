package image

import (
	"MelkOnline/internal/infrastructure"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type ImageRepository struct {
	DBconn *gorm.DB
}

func NewImageRepository() *ImageRepository {
	return &ImageRepository{
		DBconn: infrastructure.GetDB()}
}

func (ir *ImageRepository) StoreImage(file *multipart.FileHeader, AD_ID int) error {
	image, err := file.Open()
	if err != nil {
		return err
	}
	defer image.Close()
	imageFileName := filepath.Join("images", file.Filename)
	dst, err := os.Create(imageFileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, image); err != nil {
		return err
	}
	err = ir.DBconn.Exec("INSERT INTO images (AD_ID, path) VALUES (?, ?)", AD_ID, imageFileName).Error
	if err != nil {
		return err
	}
	return nil
}
