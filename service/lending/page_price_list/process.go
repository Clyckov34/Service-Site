package page_price_list

import "database/sql"

type Services struct {
	Id       int           `db:"service.id"`
	Title    string        `db:"service.title"`
	Price    int           `db:"service.price"`
	Sale     sql.NullInt64 `db:"service.sale"`
	FileName string        `db:"service.file_name"`
	Url      string        `db:"service_name.url"`
}
