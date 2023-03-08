package page_task

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
	fl "repair/pkg/file"
)

// DeleteTask удаляет задачу
func (m *Param) DeleteTask() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Delete("task").Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// DeleteFile Удаляет файл
func (m *Param) DeleteFile(nameFile []sql.NullString, dir string) error {
	for _, v := range nameFile {
		_ = fl.Delete(v.String, dir)
	}

	return nil
}
