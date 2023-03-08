package page_reg

import (
	"github.com/doug-martin/goqu/v9"
)

type SqlTX struct {
	Tx *goqu.TxDatabase
}

type Manager struct {
	Login     string
	Password  string
	FirstName string
	Email     string
	Admin     bool
}