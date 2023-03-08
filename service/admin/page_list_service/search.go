package page_list_service

import (
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// Search Поиск услуг
func (m *Param) Search() ([]ServicePriceList, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]ServicePriceList, 0)
	err = dialect.Select(&ServicePriceList{}).From("service").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service.id_name": goqu.I("service_name.id")})).Where(goqu.I("service.title").ILike("%" + m.Title + "%")).Order(goqu.I("service.title").Asc()).ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
