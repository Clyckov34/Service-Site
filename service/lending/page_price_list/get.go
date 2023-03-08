package page_price_list

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// GetServiceAll выгрузка всех услуг
func GetServiceAll() ([]Services, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Services, 0)
	err = dialect.Select(&Services{}).From("service").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service.id_name": goqu.I("service_name.id")})).Order(goqu.I("service.title").Desc()).ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetServiceName поиск по конкретным услугам
func GetServiceName(search string) ([]Services, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Services, 0)
	err = dialect.Select(&Services{}).From("service_name").Join(goqu.T("service"), goqu.On(goqu.Ex{"service.id_name": goqu.I("service_name.id")})).Order(goqu.I("service.title").Desc()).Where(goqu.I("service.title").ILike("%" + search + "%")).ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
