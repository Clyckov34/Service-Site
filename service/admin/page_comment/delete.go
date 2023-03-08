package page_comment

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

//DeleteComment Удаляет комментарий из бд
func (m *Message) DeleteComment() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Delete("comment").Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

//DeleteNameFileDB удаляет запись о фото в бд
func (m *Message) DeleteNameFileDB() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("comment").Set(goqu.Record{"file_name": nil}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
