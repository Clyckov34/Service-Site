package image

import (
	"github.com/disintegration/imaging"
)

//ResizeCenter сжатие изображения по центру
func ResizeCenter(nameFile string, width, height int) error {

	src, err := imaging.Open(nameFile)
	if err != nil {
		return err
	}

	src = imaging.Fill(src, width, height, imaging.Center, imaging.Lanczos)
	img := imaging.Blur(src, 0)

	if err := imaging.Save(img, nameFile); err != nil {
		return err
	}

	return nil
}

//Resize сжатие изображения
func Resize(nameFile string, width, height int) error {

	src, err := imaging.Open(nameFile)
	if err != nil {
		return err
	}

	src = imaging.Resize(src, width, height, imaging.Lanczos)
	img := imaging.Blur(src, 0)

	if err := imaging.Save(img, nameFile); err != nil {
		return err
	}

	return nil
}
