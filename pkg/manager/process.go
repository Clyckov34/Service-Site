package manager

import (
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

// GetEmailTypeAdmin получить список email с доступом Admin
func GetEmailTypeAdmin() ([]string, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]string, 0)
	err = dialect.Select("email").From("auth").Where(goqu.I("admin").In(true)).ScanVals(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
