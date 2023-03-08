package page_setting

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
	"repair/pkg/hesh"
)

// EditEmail редактирует почту
func (m *Param) EditEmail() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("auth").Set(goqu.Record{"email": m.Email}).Where(goqu.C("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// EditFullName изменяет ФИО
func (m *Param) EditFullName() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("auth").Set(goqu.Record{"first_name": m.FullName}).Where(goqu.C("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// EditLogin изменяет логин
func (m *Param) EditLogin() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return err
	}

	var login string
	ok, err := dialect.Select("login").From("auth").Where(goqu.C("login").In(m.Login)).Limit(1).ScanVal(&login)
	if err != nil {
		tx.Rollback()
		return err
	}

	if ok {
		tx.Rollback()
		return errors.New("менеджер с таким логином существует: ")
	}

	_, err = tx.Update("auth").Set(goqu.Record{"login": m.Login}).Where(goqu.C("id").In(m.Id)).Executor().Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// EditCheckPassword проверяет текущий пароль, и редактирует создает пароль
func (m *Param) EditCheckPassword(nowPassword, newPassword string) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return err
	}

	var password string
	_, err = tx.Select("password").From("auth").Where(goqu.C("id").In(m.Id)).Limit(1).ScanVal(&password)
	if err != nil {
		tx.Rollback()
		return err
	}

	if !hesh.Check(nowPassword, password) {
		tx.Rollback()
		return errors.New("неверный пароль")
	}

	_, err = tx.Update("auth").Set(goqu.Record{"password": hesh.NewHesh512(newPassword)}).Where(goqu.C("id").In(m.Id)).Executor().Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// EditPassword редактирует пароль
func (m *Param) EditPassword(newPassword string) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("auth").Set(goqu.Record{"password": hesh.NewHesh512(newPassword)}).Where(goqu.C("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
