package page_store

import (
	"errors"
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

// GetServiceName выгрузка конкретных услуг
func GetServiceName(id int) ([]Service, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Service, 0)
	err = dialect.Select(&Service{}).From("service").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service.id_name": goqu.I("service_name.id")})).Where(goqu.I("service_name.id").In(id)).ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}


// GetServiceID выгрузка конкретных услуг
func GetServiceID(category string, idService int) (Store, error) {
	db, err := database.Open()
	if err != nil {
		return Store{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = Store{}
	ok, err := dialect.Select(&Store{}).From("service_name").Join(goqu.T("service"), goqu.On(goqu.Ex{"service_name.id": goqu.I("service.id_name")})).Where(goqu.I("service_name.url").In(category), goqu.I("service.id").In(idService)).ScanStruct(&data)
	if err != nil {
		return Store{}, err
	}

	if !ok {
		return Store{}, errors.New("неверный ID сервиса")
	}

	return data, nil
}