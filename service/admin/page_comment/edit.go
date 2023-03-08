package page_comment

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

//Edit редактирует комментарий
func (m *Message) Edit() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("comment").Set(goqu.Record{"text": m.Text}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

//EditNameFileDB редактирует фото
func (m *Comment) EditNameFileDB(nameFile string) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("comment").Set(goqu.Record{"file_name": nameFile}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
