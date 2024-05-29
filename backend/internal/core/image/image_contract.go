package image

import "mime/multipart"

type ImageInfrasctructureContract interface {
	StoreImage(file multipart.FileHeader, AD_ID int) error
	LoadImage(AD_ID int) (string, error)
}
