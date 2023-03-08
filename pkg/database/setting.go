package database

import (
	"os"
)

type params struct {
	driver   string
	host     string
	port     string
	login    string
	password string
	dbName   string
}

// options() параметры подключения
func options() params {
	return params{
		driver:   "mysql",
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		login:    os.Getenv("DB_LOGIN"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
	}
}
