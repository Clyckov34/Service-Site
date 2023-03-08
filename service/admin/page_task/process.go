package page_task

import "database/sql"

type Param struct {
	Id                 int
	IdTask             int
	FirstNameManagerId int
	DateTime           string
	Price              int
	Address            string
	Email              string
	Phone              string
	FirstName          string
}

// Task Задачи
type Task struct {
	Id         int    `db:"task.id"`
	KeyType    string `db:"service_name.key_type"`
	DateStatus string `db:"task.date_status"`
}

type TaskChan struct {
	Data  []Task
	Error error
}

// TaskName Подробно о задачке
type TaskName struct {
	Category         string         `db:"service_name.title"`
	KeyType          string         `db:"service_name.key_type"`
	Title            string         `db:"service.title"`
	FirstNameManager sql.NullString `db:"auth.first_name"`
	FirstName        string         `db:"task.first_name"`
	Phone            string         `db:"task.phone"`
	Email            sql.NullString `db:"task.email"`
	Address          sql.NullString `db:"task.address"`
	PriceWork        sql.NullInt64  `db:"task.price"`
	Price			 int			`db:"service.price"`
	Sale             sql.NullInt64  `db:"service.sale"`
	FileName         string         `db:"service.file_name"`
	Status           string         `db:"status.name"`
	StatusTranslate  string         `db:"status.translate"`
	DateStart        string         `db:"task.date_start"`
	DateStatus       string         `db:"task.date_status"`
}

type TaskNameChan struct {
	Data  TaskName
	Error error
}

// TaskNameId Выгрузка конкретной задачи
type TaskNameId struct {
	FirstName string         `db:"task.first_name"`
	Phone     string         `db:"task.phone"`
	Email     sql.NullString `db:"task.email"`
	Address   sql.NullString `db:"task.address"`
	Price     sql.NullInt64  `db:"task.price"`
	KeyType   string         `db:"service_name.key_type"`
}

// Status Статусы
type Status struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Translate string `db:"translate"`
}

type StatusChan struct {
	Data  []Status
	Error error
}

// Comment Комментарии
type Comment struct {
	Id        int            `db:"comment.id"`
	Text      string         `db:"comment.text"`
	FirstName sql.NullString `db:"auth.first_name"`
	FileName  sql.NullString `db:"comment.file_name"`
	Date      sql.NullString `db:"comment.date"`
}

type CommentChan struct {
	Data  []Comment
	Error error
}

type Manager struct {
	Id        int    `db:"id"`
	FirstName string `db:"first_name"`
}

type ManagerChan struct {
	Data  []Manager
	Error error
}

type TaskHistory struct {
	FirstName sql.NullString `db:"auth.first_name"`
	Title     string         `db:"history_task.title"`
	Date      sql.NullString `db:"history_task.date"`
	Ip        string         `db:"history_task.ip"`
}

type TaskHistoryChan struct {
	Data  []TaskHistory
	Error error
}
