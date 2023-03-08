package page_static

import "github.com/doug-martin/goqu/v9"

type Begin struct {
	Tx      *goqu.TxDatabase
	DateNow string
}

type Status struct {
	ToDo       int64
	InProgress int64
	Pause      int64
	Denied     int64
	Done       int64
}

type StatusChan struct {
	Data Status
	Err    error
}
