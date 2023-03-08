package page_list_service

import "database/sql"

type Param struct {
	Id       int
	Title    string
	FileName string
	Text     string
	Price    int
	Sale     int
	Type     int
}

type TypeRepair struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
	Type  string `db:"key_type"`
}

type TypeRepairData struct {
	Data  []TypeRepair
	Error error
}

type ServicePriceList struct {
	Id              int           `db:"service.id"`
	TitleTypeRepair string        `db:"service_name.title"`
	Title           string        `db:"service.title"`
	Price           int           `db:"service.price"`
	Sale            sql.NullInt64 `db:"service.sale"`
}

type ServicePriceListId struct {
	Id       int           `db:"service.id"`
	Title    string        `db:"service.title"`
	Price    int           `db:"service.price"`
	Sale     sql.NullInt64 `db:"service.sale"`
	FileName string        `db:"service.file_name"`
	Text     string        `db:"service.text"`
}

type ServicePriceListData struct {
	Data  []ServicePriceList
	Error error
}
