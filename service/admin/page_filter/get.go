package page_filter

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
	"time"
)

var timeFormat = "2006-01-02"

// GetStatus получить список статусов
func GetStatus(ch chan StatusChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- StatusChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var status = make([]Status, 0)
	err = dialect.Select(&Status{}).From("status").ScanStructs(&status)
	if err != nil {
		ch <- StatusChan{nil, err}
		return
	}

	ch <- StatusChan{status, nil}
}

// GetCategory получить список категориев
func GetCategory(ch chan CategoryChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- CategoryChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var category = make([]Category, 0)
	err = dialect.Select(&Category{}).From("service_name").ScanStructs(&category)
	if err != nil {
		ch <- CategoryChan{nil, err}
		return
	}

	ch <- CategoryChan{category, nil}
}

// GetTaskFilter выгрузка задач по фильтрам
func (m *Filter) GetTaskFilter() ([]Tasks, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var tasks = make([]Tasks, 0)
	err = dialect.Select(&Tasks{}).From("task").LeftJoin(goqu.T("auth"), goqu.On(goqu.Ex{"auth.id": goqu.I("task.id_auth")})).
		Join(goqu.T("status"), goqu.On(goqu.Ex{"status.id": goqu.I("task.id_status")})).
		Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service_name.id": goqu.I("task.id_type")})).
		Where(goqu.I("task.id").ILike("%"+m.Task.Number+"%"), goqu.I("service_name.key_type").ILike("%"+m.Task.Key+"%"), goqu.I("status.id").ILike("%"+m.Status+"%"), goqu.I("service_name.id").ILike("%"+m.Category+"%"),
			goqu.I("task.first_name").ILike("%"+m.FirstName+"%"), goqu.I("task.phone").ILike("%"+m.Phone+"%") ,goqu.COALESCE(goqu.I("task.address"), "").ILike("%"+m.Address+"%"), goqu.COALESCE(goqu.I("auth.id"), "").ILike("%"+m.Manager+"%"),
			goqu.I("task.date_status").Gte(m.DateStart), goqu.I("task.date_status").Lte(addDate(m.DateEnd, 0, 0, 1))).Order(goqu.I("task.date_status").Desc()).Distinct().ScanStructs(&tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// addDate добавляет интервал в дату
func addDate(nowDate string, years, months, days int) string {
	h, _ := time.Parse(timeFormat, nowDate)
	return h.AddDate(years, months, days).Format(timeFormat)
}

