package page_type_service

import (
	"repair/pkg/database"
	"strings"

	"github.com/doug-martin/goqu/v9"
)

// CreateService загрузка профиля
func (m *Param) CreateService(nameFile string) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Insert("service_name").Rows(goqu.Record{"key_type": strings.ToUpper(m.KeyType), "title": m.Title, "url": m.Url, "file_name": nameFile}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
