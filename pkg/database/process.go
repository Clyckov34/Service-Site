package database

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// Open подключения к бд
func Open() (*sql.DB, error) {
	param := options()

	db, err := sql.Open(param.driver, fmt.Sprintf("%v:%v@tcp(%v%v)/%v", param.login, param.password, param.host, param.port, param.dbName))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Dialect обработчик sql генератор
func Dialect(db *sql.DB) *goqu.Database {
	param := options()
	return goqu.Dialect(param.driver).DB(db)
}
