package page_list_service

import (
	"repair/pkg/check"
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

// Add добавить услугу в БД
func (m *Param) Add() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Insert("service").Rows(goqu.Record{"id_name": m.Type, "title": m.Title, "price": m.Price, "sale": check.NullIntBD(m.Sale), "text": m.Text, "file_name": m.FileName}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
