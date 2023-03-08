package page_comment

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/check"
	"repair/pkg/database"
)

// AddBD добавляет запись в бд
func (m *Comment) AddBD(newNameFile string) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)
	_, err = dialect.Insert("comment").Rows(goqu.Record{"id_task": m.Id, "id_auth": m.IdManager, "text": m.Text, "file_name": check.NullStringBD(newNameFile), "date": m.Date}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
