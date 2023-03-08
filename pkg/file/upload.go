package file

import (
	"mime/multipart"
	"path/filepath"
	"repair/pkg/check"
	"repair/pkg/image"
	"time"

	"github.com/gin-gonic/gin"
)

var nowDate = time.Now().Format("2006-01-02 15:04:05")
const (
	widht = 700
	height = 500
)

// UploadAndCompress загрузка файла и сжатие
func UploadAndCompress(ctx *gin.Context, file *multipart.FileHeader, dir string) (nameFile string, err error) {
	fileName := nowDate + " " + filepath.Base(file.Filename)

	//проверка тип файла
	if err := check.TypeFile(fileName, ".png", ".jpeg", ".jpg"); err != nil {
		return "", err
	}

	//загрузка файла на сервер
	if err := ctx.SaveUploadedFile(file, dir+fileName); err != nil {
		return "", err
	}

	//сжатие изображения
	if err := image.Resize(dir+fileName, widht, height); err != nil {
		return "", err
	}

	return fileName, nil
}

// UploadAllAndCompress загрузка несколько файлов и сжатие
func UploadAllAndCompress(ctx *gin.Context, files []*multipart.FileHeader, dir string) error {
	for _, file := range files {
		fileName := nowDate + " " + filepath.Base(file.Filename)

		//проверка тип файла
		if err := check.TypeFile(fileName, ".png", ".jpeg", ".jpg"); err != nil {
			return err
		}

		//загрузка файла на сервер
		if err := ctx.SaveUploadedFile(file, dir+fileName); err != nil {
			return err
		}

		//сжатие изображения
		if err := image.Resize(dir+fileName, widht, height); err != nil {
			return err
		}

	}
	return nil
}