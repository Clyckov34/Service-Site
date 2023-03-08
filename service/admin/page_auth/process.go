package page_auth

import (
	"github.com/doug-martin/goqu/v9"
)

type Page struct {
	Login    string
	Password string
	Email    string
}

type Data struct {
	Manager  Manager
	Security Security
}

type Manager struct {
	Id       int    `db:"id"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Admin    bool   `db:"admin"`
}

type Security struct {
	Code     string
	DateTime string
}

type auth struct {
	Tx *goqu.TxDatabase
}

type Check struct {
	Id       int    `db:"id"`
	Code     string `db:"code"`
	DateTime string `db:"date"`
}
