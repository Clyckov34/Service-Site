package page_manager

import "database/sql"

type Param struct {
	Id   int
	Type sql.NullBool
}

type Manager struct {
	Id        int          `db:"id"`
	Login     string       `db:"login"`
	FirstName string       `db:"first_name"`
	Type      sql.NullBool `db:"admin"`
}
