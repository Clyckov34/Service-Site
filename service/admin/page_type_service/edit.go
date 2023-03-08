package page_type_service

import (
	"repair/pkg/database"
	"strings"

	"github.com/doug-martin/goqu/v9"
)

// Edit редактирует данные
func (m *Param) Edit() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("service_name").Set(goqu.Record{"title": m.Title, "url": m.Url, "key_type": strings.ToUpper(m.KeyType)}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// UpdatePortfolio обновляет данные в профиле
func (m *Param) UpdatePortfolio(nameFile string) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("service_name").Set(goqu.Record{"file_name": nameFile}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return err
}
