package page_auth

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
	"repair/pkg/generation"
	"repair/pkg/hesh"
)

// AuthCabinet авторизация в личном кабинете
func (m *Page) AuthCabinet() (Data, error) {
	db, err := database.Open()
	if err != nil {
		return Data{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return Data{}, err
	}

	manager := auth{Tx: tx}

	//Авторизация
	mng, err := manager.auth(m.Login, m.Password)
	if err != nil {
		tx.Rollback()
		return Data{}, err
	}

	//Генерация секретного кода + интервал времяни
	sec, err := manager.security(mng.Id)
	if err != nil {
		tx.Rollback()
		return Data{}, err
	}

	if err := tx.Commit(); err != nil {
		return Data{}, err
	}

	return Data{
		Manager: Manager{
			Id:    mng.Id,
			Email: mng.Email,
			Admin: mng.Admin,
		},
		Security: Security{
			Code:     sec.Code,
			DateTime: sec.DateTime,
		},
	}, nil

}

// auth авторизация пользователя
func (m *auth) auth(login, password string) (Manager, error) {

	var mng = Manager{}
	ok, err := m.Tx.Select(&Manager{}).From("auth").Where(goqu.I("login").In(login)).Limit(1).ScanStruct(&mng)
	if err != nil {
		return Manager{}, err
	}

	if !ok {
		return Manager{}, errors.New("нет такого администратора")
	}

	if !hesh.Check(password, mng.Password) {
		return Manager{}, errors.New("неверный пароль")
	}

	return mng, nil
}

// security генерация даты на интервал времяни + секретный код
func (m *auth) security(id int) (Security, error) {
	dateChan := make(chan string)
	codeChan := make(chan string)

	go generation.DateTimeAddChan(10, dateChan)
	go generation.CodeChan(6, codeChan)

	date := <-dateChan
	code := <-codeChan

	_, err := m.Tx.Update("auth").Set(goqu.Record{"date": date, "code": hesh.NewHesh512(code)}).Where(goqu.I("id").In(id)).Executor().Exec()
	if err != nil {
		return Security{}, err
	}

	return Security{Code: code, DateTime: date}, nil
}
