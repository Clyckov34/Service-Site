package page_social_network

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// Edit редактирует данные
func (m *Param) Edit() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("social_network").Set(goqu.Record{"id_icon": m.IdIcon, "url": m.Url}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
