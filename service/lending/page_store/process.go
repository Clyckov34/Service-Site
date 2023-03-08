package page_store

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Service struct {
	Id       int           `db:"service.id"`
	Title    string        `db:"service.title"`
	Price    int           `db:"service.price"`
	Sale     sql.NullInt64 `db:"service.sale"`
	FileName string        `db:"service.file_name"`
	Url      string        `db:"service_name.url"`
}

type Store struct {
	ServiceId       int           `db:"service.id"`
	ServiceTitle    string        `db:"service.title"`
	ServicePrice    int           `db:"service.price"`
	ServiceSale     sql.NullInt64 `db:"service.sale"`
	ServiceFileName string        `db:"service.file_name"`
	ServiceText     string        `db:"service.text"`
	CategoryID      int           `db:"service_name.id"`
	CategoryTitle   string        `db:"service_name.title"`
}

type StoreDB struct {
	Category string `db:"service_name.title"`
	Service  string `db:"service.title"`
	FileName string `db:"service.file_name"`
}

type Send struct {
	Street     string
	IdService  int
	IdCategory int
	FirstName  string
	Phone      string
	Email      string
	Date       string
}

type Task struct {
	Tx *goqu.TxDatabase
}
