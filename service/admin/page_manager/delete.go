package page_manager

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

//Delete уадялет менеджера
func (m *Param) Delete() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Delete("auth").Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
