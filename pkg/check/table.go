package check

import (
	"repair/pkg/database"
)

// InstallTable проверка на существеющие таблицы
func InstallTable(nameTable string) (int64, error) {
	db, err := database.Open()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	id, err := dialect.From(nameTable).Count()
	if err != nil {
		return 0, err
	}

	return id, err
}
