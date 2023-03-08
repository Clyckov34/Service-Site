package page_cabinet

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// CheckTask проверка задач на статус
func CheckTask(nameStatus string) (bool, error) {
	db, err := database.Open()
	if err != nil {
		return false, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var task int
	ok, err := dialect.Select("task.id").From("task").Join(goqu.T("status"), goqu.On(goqu.Ex{"task.id_status": goqu.I("status.id")})).Where(goqu.I("status.name").In(nameStatus)).Limit(1).ScanVal(&task)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, errors.New("ошибка нет задач")
	}

	return true, nil
}
