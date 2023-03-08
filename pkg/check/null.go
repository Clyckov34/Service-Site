package check

import (
	"database/sql"
	"strings"
)

// NullStringBD проверка БД на пустату строка
func NullStringBD(data string) sql.NullString {
	data = strings.TrimSpace(data)

	if len(data) != 0 {
		return sql.NullString{String: data, Valid: true}
	} else {
		return sql.NullString{}
	}
}

// NullIntBD проверка БД на пустату целое число
func NullIntBD(number int) sql.NullInt64 {
	if number > 0 {
		return sql.NullInt64{Int64: int64(number), Valid: true}
	} else {
		return sql.NullInt64{}
	}
}
