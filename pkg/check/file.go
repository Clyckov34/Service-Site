package check

import (
	"errors"
	"path/filepath"
)

//TypeFile Проверка на тип файла
func TypeFile(nameFile string, typeFile ...string) error {
	extension := filepath.Ext(nameFile)

	for _, tf := range typeFile {
		if extension != tf {
			continue
		}

		return nil
	}

	return errors.New("неверный формат файла")
}
