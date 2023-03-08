package page_list_service

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// GetTypeRepair получить ремонт
func GetTypeRepair(ch chan TypeRepairData) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- TypeRepairData{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]TypeRepair, 0)
	err = dialect.Select(&TypeRepair{}).From("service_name").ScanStructs(&data)
	if err != nil {
		ch <- TypeRepairData{nil, err}
		return
	}

	ch <- TypeRepairData{Data: data, Error: nil}
}

// GetPriceListService Выгрузка прайс-лист услуги
func GetPriceListService(ch chan ServicePriceListData) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- ServicePriceListData{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]ServicePriceList, 0)
	err = dialect.Select(&ServicePriceList{}).From("service").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service.id_name": goqu.I("service_name.id")})).ScanStructs(&data)
	if err != nil {
		ch <- ServicePriceListData{nil, err}
		return
	}

	ch <- ServicePriceListData{Data: data, Error: nil}
}

// GetServiceId выгрузка конкретного сервиса
func (m *Param) GetServiceId() (ServicePriceListId, error) {
	db, err := database.Open()
	if err != nil {
		return ServicePriceListId{}, nil
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = ServicePriceListId{}
	ok, err := dialect.Select(&ServicePriceListId{}).From("service").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service.id_name": goqu.I("service_name.id")})).Where(goqu.I("service.id").In(m.Id)).ScanStruct(&data)
	if err != nil {
		return ServicePriceListId{}, nil
	}

	if !ok {
		return ServicePriceListId{}, errors.New("ошибка: неверный ID")
	}

	return data, nil
}
