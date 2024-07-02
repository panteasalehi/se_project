package ADregister

import (
	"mime/multipart"

	"gorm.io/gorm"

	model "MelkOnline/internal/core"
	"MelkOnline/internal/infrastructure"
	"MelkOnline/internal/infrastructure/image"
)

type ADregisterRepository struct {
	DBconn *gorm.DB
	ir     *image.ImageRepository
}

func NewADregisterRepository() *ADregisterRepository {
	return &ADregisterRepository{DBconn: infrastructure.GetDB(), ir: image.NewImageRepository()}
}

func (sr *ADregisterRepository) StoreAD(Ad *model.AD, image *multipart.FileHeader) (int, error) {
	err := sr.DBconn.Create(&Ad).Error
	if err != nil {
		return 0, err
	}
	err = sr.ir.StoreImage(image, Ad.ID)
	if err != nil {
		sr.DBconn.Delete(&Ad)
		return 0, err
	}
	return Ad.ID, nil
}
