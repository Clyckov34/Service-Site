package page_list_service

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

//Delete удаляет запись о услуги
func (m *Param) Delete() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Delete("service").Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
