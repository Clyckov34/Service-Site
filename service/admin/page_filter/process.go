package page_filter

import "database/sql"

type Filter struct {
	Status    string
	Category  string
	FirstName string
	Phone     string
	Task      Parser
	Address   string
	Manager   string
	DateStart string
	DateEnd   string
}

type Tasks struct {
	Id              int            `db:"task.id"`
	Category        string         `db:"service_name.title"`
	KeyType         string         `db:"service_name.key_type"`
	Status          string         `db:"status.name"`
	StatusTranslate string         `db:"status.translate"`
	FirstName       string         `db:"task.first_name"`
	Phone           string         `db:"task.phone"`
	Address         sql.NullString `db:"task.address"`
	Manager         sql.NullString `db:"auth.first_name"`
	DateStart       string         `db:"task.date_start"`
	DateStatus      string         `db:"task.date_status"`
}

type Status struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Translate string `db:"translate"`
}

type StatusChan struct {
	Data  []Status
	Error error
}

type Category struct {
	Id   int    `db:"id"`
	Name string `db:"title"`
}

type CategoryChan struct {
	Data  []Category
	Error error
}

type Parser struct {
	Key    string
	Number string
}
