package page_setting

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// GetSetting получить данные о авторезованным пользователям
func (m *Param) GetSetting() (Param, error) {
	db, err := database.Open()
	if err != nil {
		return Param{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = Param{}
	ok, err := dialect.Select(&Param{}).From("auth").Where(goqu.I("id").In(m.Id)).ScanStruct(&data)
	if err != nil || !ok {
		return Param{}, errors.New("ошибка выгрузки настроек: " + err.Error())
	}

	return data, nil
}
