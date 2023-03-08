package page_task

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// GetTaskName Выгрузка задач конкретного статуса
func GetTaskName(statusName string, limit uint, ch chan TaskChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- TaskChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	data := make([]Task, 0)
	err = dialect.Select(&Task{}).From("task").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"task.id_type": goqu.I("service_name.id")})).Join(goqu.T("status"), goqu.On(goqu.Ex{"task.id_status": goqu.I("status.id")})).Where(goqu.I("status.name").In(statusName)).Order(goqu.I("task.date_status").Desc()).Limit(limit).ScanStructs(&data)
	if err != nil {
		ch <- TaskChan{nil, err}
		return
	}

	ch <- TaskChan{data, nil}
}

// GetStatus выводит все статусы
func GetStatus(ch chan StatusChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- StatusChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	data := make([]Status, 0)
	err = dialect.Select(&Status{}).From("status").Where(goqu.I("name").NotIn("To Do")).ScanStructs(&data)
	if err != nil {
		ch <- StatusChan{nil, err}
		return
	}

	ch <- StatusChan{data, nil}
}

// GetTaskNameDetail выгрузка конкретной подробной задачи
func (m *Param) GetTaskNameDetail(ch chan TaskNameChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- TaskNameChan{TaskName{}, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = TaskName{}
	ok, err := dialect.Select(&TaskName{}).From("task").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"task.id_type": goqu.I("service_name.id")})).
		Join(goqu.T("status"), goqu.On(goqu.Ex{"task.id_status": goqu.I("status.id")})).Join(goqu.T("service"), goqu.On(goqu.Ex{"service.id": goqu.I("task.id_service")})).
		LeftJoin(goqu.T("auth"), goqu.On(goqu.Ex{"task.id_auth": goqu.I("auth.id")})).Where(goqu.I("task.id").In(m.Id)).ScanStruct(&data)

	if err != nil {
		ch <- TaskNameChan{TaskName{}, err}
		return
	}

	if !ok {
		ch <- TaskNameChan{TaskName{}, errors.New("неверный ID задачи")}
		return
	}

	ch <- TaskNameChan{data, nil}

}

// GetTaskNameDetailId выгрузка конкретной подробной задачи
func (m *Param) GetTaskNameDetailId() (TaskNameId, error) {
	db, err := database.Open()
	if err != nil {
		return TaskNameId{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = TaskNameId{}
	ok, err := dialect.Select(&TaskNameId{}).From("task").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"task.id_type": goqu.I("service_name.id")})).Where(goqu.I("task.id").In(m.Id)).ScanStruct(&data)
	if err != nil {
		return TaskNameId{}, err
	}

	if !ok {
		return TaskNameId{}, errors.New("ошибка ID задачи")
	}

	return data, nil
}

// GetComment выгрузка комментарий
func (m *Param) GetComment(ch chan CommentChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- CommentChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Comment, 0)
	err = dialect.Select(&Comment{}).From("comment").Join(goqu.T("task"), goqu.On(goqu.Ex{"comment.id_task": goqu.I("task.id")})).
		LeftJoin(goqu.T("auth"), goqu.On(goqu.Ex{"comment.id_auth": goqu.I("auth.id")})).Where(goqu.I("task.id").In(m.Id)).
		Order(goqu.I("comment.date").Desc()).ScanStructs(&data)

	if err != nil {
		ch <- CommentChan{nil, err}
		return
	}

	ch <- CommentChan{data, nil}
}

// GetManager выгрузка менеджера
func (m *Param) GetManager(ch chan ManagerChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- ManagerChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Manager, 0)
	err = dialect.Select(&Manager{}).From("auth").ScanStructs(&data)
	if err != nil {
		ch <- ManagerChan{nil, err}
		return
	}

	ch <- ManagerChan{data, nil}
}

//GetTaskHistory выгрузка истории задач
func (m *Param) GetTaskHistory(ch chan TaskHistoryChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- TaskHistoryChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]TaskHistory, 0)
	err = dialect.Select(&TaskHistory{}).From("history_task").LeftJoin(goqu.T("auth"), goqu.On(goqu.Ex{"history_task.id_auth": goqu.I("auth.id")})).Where(goqu.I("history_task.id_task").In(m.Id)).Order(goqu.I("history_task.date").Desc()).ScanStructs(&data)
	if err != nil {
		ch <- TaskHistoryChan{nil, err}
		return
	}

	ch <- TaskHistoryChan{data, nil}
}