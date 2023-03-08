package page_auth

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
	"repair/pkg/hesh"
)

// CodeAndDateTime проверка секретного кода и время
func (m *Check) CodeAndDateTime() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var user = Check{}
	ok, err := dialect.Select(&Check{}).From("auth").Where(goqu.I("id").In(m.Id)).ScanStruct(&user)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("ошибка проверки данных")
	}

	// Проверка времяни
	if err = m.timeout(user.DateTime); err != nil {
		return err
	}

	// Проверка кода
	if !m.code(user.Code) {
		return errors.New("код не совпадает")
	}

	return nil
}

// timeout проверка времяни
func (m *Check) timeout(date string) error {
	if m.DateTime >= date {
		return errors.New("превысило время ожидание")
	} else {
		return nil
	}
}

// code проверка одноразового кода
func (m *Check) code(heshCode string) bool {
	return hesh.Check(m.Code, heshCode)
}
