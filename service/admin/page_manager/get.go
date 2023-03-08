package page_manager

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// GetManager получить список менеджеров, кроме активных
func (m *Param) GetManager() ([]Manager, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Manager, 0)
	err = dialect.Select(&Manager{}).From("auth").Where(goqu.I("id").NotIn(m.Id)).ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetManagerName Выгрузка конкретного менеджера
func (m *Param) GetManagerName() (string, error) {
	db, err := database.Open()
	if err != nil {
		return "", err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var firstName string
	ok, err := dialect.Select("first_name").From("auth").Where(goqu.I("id").In(m.Id)).ScanVal(&firstName)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("неверный ID менеджера")
	}

	return firstName, nil
}
