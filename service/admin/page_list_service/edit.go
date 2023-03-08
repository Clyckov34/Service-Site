package page_list_service

import (
	"repair/pkg/check"
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

// Edit Изменения данных
func (m *Param) Edit() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("service").Set(goqu.Record{"title": m.Title, "price": m.Price, "sale": check.NullIntBD(m.Sale), "text": m.Text}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

//UpdatePortfolio редактирует запись о файле
func (m *Param) UpdatePortfolio(nameFile string) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("service").Set(goqu.Record{"file_name": nameFile}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return err
}