package page_social_network

import (
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// GetAll вывод данных список соц сетей
func (m *Param) GetAll(ch chan ListNetworkAllChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- ListNetworkAllChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]ListNetworkAll, 0)
	err = dialect.Select(&ListNetworkAll{}).From("social_network").Join(goqu.T("icon"), goqu.On(goqu.Ex{"social_network.id_icon": goqu.I("icon.id")})).ScanStructs(&data)
	if err != nil {
		ch <- ListNetworkAllChan{nil, err}
		return
	}

	ch <- ListNetworkAllChan{data, nil}
}

// GetIcon получить иконку
func (m *Param) GetIcon(ch chan IconChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- IconChan{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Icon, 0)
	err = dialect.Select(&Icon{}).From("icon").ScanStructs(&data)
	if err != nil {
		ch <- IconChan{nil, err}
		return
	}

	ch <- IconChan{data, nil}
}

// GetID Выводит конкретный ID соц.сети
func (m *Param) GetID(ch chan ListNetworkIdChan) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- ListNetworkIdChan{ListNetworkId{}, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = ListNetworkId{}
	ok, err := dialect.Select(&ListNetworkId{}).From("social_network").Where(goqu.I("id").In(m.Id)).ScanStruct(&data)
	if err != nil {
		ch <- ListNetworkIdChan{ListNetworkId{}, err}
		return
	}

	if !ok {
		ch <- ListNetworkIdChan{ListNetworkId{}, errors.New("ошибка: неверный ID")}
		return
	}

	ch <- ListNetworkIdChan{data, nil}
}
