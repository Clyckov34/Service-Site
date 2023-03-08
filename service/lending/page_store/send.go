package page_store

import (
	"errors"
	"log"
	"repair/internal/page/admin/page_list_service"
	"repair/pkg/check"
	"repair/pkg/database"
	ml "repair/pkg/mail"
	"repair/pkg/manager"

	"github.com/doug-martin/goqu/v9"
)

// SendMail отправка на почту
func (m *Send) SendMail(st *StoreDB) error {
	listEmail, err := manager.GetEmailTypeAdmin()
	if err != nil {
		return err
	}

	sendMessage := func(email string) {
		fileName := page_list_service.Dir + st.FileName
		if err := ml.Send(email, "Заявка на оформления услуги", message(m.FirstName, m.Phone, m.Email, m.Street, st.Category, st.Service), fileName); err != nil {
			log.Println(err)
		}
	}

	for _, v := range listEmail {
		go sendMessage(v)
	}

	return nil
}

// CreateTask создает задачу c определенным статусом
func (m *Send) CreateTask(statusName string) (*StoreDB, error) {
	db, err := database.Open()
	if err != nil {
		return &StoreDB{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return &StoreDB{}, err
	}

	task := Task{Tx: tx}

	statusID, err := task.getStatusID(statusName)
	if err != nil {
		tx.Rollback()
		return &StoreDB{}, err
	}

	if err := task.insertTask(m, statusID); err != nil {
		tx.Rollback()
		return &StoreDB{}, err
	}

	data, err := task.getStore(m.IdService)
	if err != nil {
		tx.Rollback()
		return &StoreDB{}, err
	}

	if err := tx.Commit(); err != nil {
		return &StoreDB{}, err
	}

	return &data, nil
}

// getStatusID получить ID статус
func (m *Task) getStatusID(statusName string) (int, error) {
	var statusID int
	ok, err := m.Tx.Select("id").From("status").Where(goqu.I("name").In(statusName)).ScanVal(&statusID)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("неверный статус")
	}

	return statusID, nil
}

// getStore получить название сервиса, категории, файлы
func (m *Task) getStore(serviceID int) (StoreDB, error) {
	store := StoreDB{}
	ok, err := m.Tx.Select(&StoreDB{}).From("service_name").Join(goqu.T("service"), goqu.On(goqu.Ex{"service_name.id": goqu.I("service.id_name")})).Where(goqu.I("service.id").In(serviceID)).ScanStruct(&store)
	if err != nil {
		return StoreDB{}, nil
	}

	if !ok {
		return StoreDB{}, errors.New("неверный ID service")
	}

	return store, nil
}

// insertTask создает запись в бд
func (m *Task) insertTask(t *Send, statusID int) error {
	_, err := m.Tx.Insert("task").Rows(goqu.Record{
		"id_status":   statusID,
		"id_type":     t.IdCategory,
		"id_service":  t.IdService,
		"first_name":  t.FirstName,
		"phone":       t.Phone,
		"email":       check.NullStringBD(t.Email),
		"address":     check.NullStringBD(t.Street),
		"date_start":  t.Date,
		"date_status": t.Date,
	}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
