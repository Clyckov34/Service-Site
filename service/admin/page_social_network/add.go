package page_social_network

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// Add добавляет в БД группу соц сети
func (m *Param) Add() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Insert("social_network").Rows(goqu.Record{"id_icon": m.IdIcon, "url": m.Url}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
