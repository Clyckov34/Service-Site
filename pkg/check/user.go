package check

import (
	"errors"
)

// ManagerAdmin проверка на тип учетной записи
func ManagerAdmin(admin bool) error {
	if !admin {
		return errors.New("доступ запрещен")
	}

	return nil
}
