package page_comment

import (
	"database/sql"
	"errors"
	"github.com/doug-martin/goqu/v9"
	"repair/pkg/database"
)

// GetCommentId получить конкретный комментарий
func GetCommentId(id int) (Message, error) {
	db, err := database.Open()
	if err != nil {
		return Message{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var message = Message{}
	ok, err := dialect.Select(&Message{}).From("comment").Where(goqu.I("id").In(id)).ScanStruct(&message)
	if err != nil {
		return Message{}, err
	}

	if !ok {
		return Message{}, errors.New("неверный ID комментария")
	}

	return message, err
}

// GetCommentPhoto получает список файлов закрепленном за камментариями
func GetCommentPhoto(idTask int) ([]sql.NullString, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var listFile = make([]sql.NullString, 0)
	err = dialect.Select("comment.file_name").From("comment").Join(goqu.T("task"), goqu.On(goqu.Ex{"task.id": goqu.I("comment.id_task")})).Where(goqu.I("task.id").In(idTask)).ScanVals(&listFile)
	if err != nil {
		return nil, err
	}

	return listFile, nil
}

// GetIdManagerComment получить id пользователя комментария
func GetIdManagerComment(idComment int) (int, error) {
	db, err := database.Open()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var id int
	ok, err := dialect.Select("id_auth").From("comment").Where(goqu.I("id").In(idComment)).ScanVal(&id)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("неверный ID комментария")
	}

	return id, nil
}
