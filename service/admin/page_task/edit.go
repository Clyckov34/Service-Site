package page_task

import (
	"errors"
	"repair/pkg/check"
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

// EditStatus изменяет статус
func (m *Param) EditStatus() (statusName string, err error) {
	db, err := database.Open()
	if err != nil {
		return "", err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return "", err
	}

	_, err = tx.Update("task").Set(goqu.Record{"id_status": m.Id, "date_status": m.DateTime}).Where(goqu.I("id").In(m.IdTask)).Executor().Exec()
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	var status string
	ok, err := tx.Select("name").From("status").Where(goqu.I("id").In(m.Id)).ScanVal(&status)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	if !ok {
		_ = tx.Rollback()
		return "", errors.New("неверный ID status")
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return status, nil
}

// EditPrice редактирует цену
func (m *Param) EditPrice() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("task").Set(goqu.Record{"price": check.NullIntBD(m.Price)}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// EditTask редактирует задачу
func (m *Param) EditTask() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Update("task").Set(goqu.Record{"first_name": m.FirstName, "phone": m.Phone, "email": check.NullStringBD(m.Email), "address": check.NullStringBD(m.Address), "price": check.NullIntBD(m.Price)}).Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// EditManager редактирует менеджер
func (m *Param) EditManager() (managerName string, err error) {
	db, err := database.Open()
	if err != nil {
		return "", err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return "", err
	}

	_, err = tx.Update("task").Set(goqu.Record{"id_auth": m.FirstNameManagerId}).Where(goqu.I("id").In(m.IdTask)).Executor().Exec()
	if err != nil {
		tx.Rollback()
		return "", err
	}

	var manager string
	ok, err := tx.Select("auth.first_name").From("task").Join(goqu.T("auth"), goqu.On(goqu.Ex{"task.id_auth": goqu.I("auth.id")})).Where(goqu.I("task.id").In(m.IdTask)).ScanVal(&manager)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if !ok {
		tx.Rollback()
		return "", errors.New("неверный ID  manager")
	}
	
	if err := tx.Commit(); err != nil {
		return "", err
	}


	return manager, nil
}
