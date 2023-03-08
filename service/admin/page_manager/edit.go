package page_manager

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// EditType изменяет тип учетной записи
func (m *Param) EditType() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("auth").Set(goqu.Record{"admin": m.Type.Bool}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
