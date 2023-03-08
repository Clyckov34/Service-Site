package page_comment

import (
	"database/sql"
)

type Comment struct {
	Id        int
	IdManager int
	Text      string
	Date      string
}

type Message struct {
	Id       int            `db:"id"`
	IdTask   int            `db:"id_task"`
	Text     string         `db:"text"`
	FileName sql.NullString `db:"file_name"`
}
