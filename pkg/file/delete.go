package file

import "os"

// Delete Удаляет файл
func Delete(nameFile string, dir string) error {
	if err := os.Remove(dir + nameFile); err != nil {
		return err
	}

	return nil
}
