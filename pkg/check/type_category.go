package check

import (
	"errors"
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

type DBService struct {
	Id    int    `db:"service_name.id"`
	Title string `db:"service_name.title"`
}


// Сategory проверка на тип категории
func Сategory(category string) (DBService, error) {
	db, err := database.Open()
	if err != nil {
		return DBService{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var p DBService
	ok, err := dialect.Select(&DBService{}).From("service_name").Where(goqu.I("service_name.url").In(category)).ScanStruct(&p)
	if err != nil {
		return DBService{}, err
	}

	if !ok {
		return DBService{}, errors.New("нет такого раздела услуги")
	}

	return p, nil
}
