package page_static

import (
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

// GetStaticsAll выгрузка всех статусов по кол-во задач за все время
func GetStaticsAll() (Status, error) {

	db, err := database.Open()
	if err != nil {
		return Status{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return Status{}, err
	}

	static := Begin{
		Tx: tx,
	}

	toDo, err := static.getQuantityAll("To Do")
	if err != nil {
		tx.Rollback()
		return Status{}, err
	}

	inProgress, err := static.getQuantityAll("In Progress")
	if err != nil {
		tx.Rollback()
		return Status{}, err
	}

	pause, err := static.getQuantityAll("Pause")
	if err != nil {
		tx.Rollback()
		return Status{}, err
	}

	denied, err := static.getQuantityAll("Denied")
	if err != nil {
		tx.Rollback()
		return Status{}, err
	}

	done, err := static.getQuantityAll("Done")
	if err != nil {
		tx.Rollback()
		return Status{}, err
	}

	if err := tx.Commit(); err != nil {
		return Status{}, err
	}

	return Status{
		ToDo:       toDo,
		InProgress: inProgress,
		Pause:      pause,
		Denied:     denied,
		Done:       done,
	}, err
}

// getQuantityAll Выгрузка кол задач за все вермя для статичтики
func (tx *Begin) getQuantityAll(nameStatus string) (int64, error) {
	quantity, err := tx.Tx.From("task").Join(goqu.T("status"), goqu.On(goqu.Ex{"status.id": goqu.I("task.id_status")})).Where(goqu.I("status.name").In(nameStatus)).Count()
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
