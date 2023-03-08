package page_reg

import (
	"repair/pkg/database"
	sl "repair/sql"
)

// CreateTable  созданиятаблиц в БД
func (m *Manager) CreateTable() error {

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

	var create = &SqlTX{tx}

	if err = create.auth(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.serviceName(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.service(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.icon(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.socialNetwork(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.status(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.task(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.comment(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = create.historyTask(); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}


	return nil
}

// serviceName Создание таблицы в БД... Виды Услуг.
func (m *SqlTX) serviceName() error {
	_, err := m.Tx.Exec(sl.CreateServiceName)
	if err != nil {
		return err
	}

	return nil
}

// service Создание таблицы в БД... Список Услуги.
func (m *SqlTX) service() error {
	_, err := m.Tx.Exec(sl.CreateService)
	if err != nil {
		return err
	}

	return nil
}

// socialNetwork Создание таблицы в БД... Группа социальных сетей
func (m *SqlTX) socialNetwork() error {
	_, err := m.Tx.Exec(sl.CreateSocialNetwork)
	if err != nil {
		return err
	}

	return nil
}

// auth Создание таблицы в БД... Авторизация
func (m *SqlTX) auth() error {
	_, err := m.Tx.Exec(sl.CreateAuth)
	if err != nil {
		return err
	}

	return nil
}

// status Создание таблицы в БД... Статусы
func (m *SqlTX) status() error {
	_, err := m.Tx.Exec(sl.CreateStatus)
	if err != nil {
		return err
	}

	return nil
}

// task Создание таблицы в БД... Задачи
func (m *SqlTX) task() error {
	_, err := m.Tx.Exec(sl.CreateTask)
	if err != nil {
		return err
	}

	return nil
}

// icon Создание таблицы в БД... Иконки
func (m *SqlTX) icon() error {
	_, err := m.Tx.Exec(sl.CreateIcon)
	if err != nil {
		return err
	}

	return nil
}

// comment Создание таблицы в БД... Задачи
func (m *SqlTX) comment() error {
	_, err := m.Tx.Exec(sl.CreateComment)
	if err != nil {
		return err
	}

	return nil
}

//historyTask Создает таблицу история задач
func (m *SqlTX) historyTask() error {
	_, err := m.Tx.Exec(sl.CreateHistoryTask)
	if err != nil {
		return err
	}

	return nil
}