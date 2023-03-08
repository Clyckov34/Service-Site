package file

import (
	"os"
)

// GetFileAll получить все файлы из папки
func GetFileAll(dir string) (files []string, err error) {
	file, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)

	for _, f := range file {
		if f.Name() == ".gitkeep" {
			continue
		}

		list = append(list, f.Name())
	}

	return list, nil
}
