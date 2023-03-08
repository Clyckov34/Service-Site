package page_type_service

import (
	"errors"
	"repair/pkg/database"

	"github.com/doug-martin/goqu/v9"
)

// GetProfile получить портфолио
func (m *Param) GetProfile() ([]Profile, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = make([]Profile, 0)
	err = dialect.Select(&Profile{}).From("service_name").ScanStructs(&data)
	if err != nil {
		return nil, err
	}

	return data, err
}

// GetProfileId получить подробные дынные о конкретной портфолио
func (m *Param) GetProfileId() (ProfileDetail, error) {
	db, err := database.Open()
	if err != nil {
		return ProfileDetail{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var data = ProfileDetail{}
	ok, err := dialect.Select(&ProfileDetail{}).From("service_name").Where(goqu.I("id").In(m.Id)).ScanStruct(&data)
	if err != nil {
		return ProfileDetail{}, err
	}

	if !ok {
		return ProfileDetail{}, errors.New("неверный ID профиля")
	}

	return data, nil
}

// getProfilePhotoAll получить список всех файлов закрепленные за задачими
func (m *Param) getProfileFile() (nameFiles FileName, err error) {
	db, err := database.Open()
	if err != nil {
		return FileName{}, err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	tx, err := dialect.Begin()
	if err != nil {
		return FileName{}, err
	}

	// Выгрузка список файлов закрепленными за комментариями
	var commentFile = make([]string, 0)
	err = tx.Select("comment.file_name").From("task").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service_name.id": goqu.I("task.id_type")})).
		Join(goqu.T("comment"), goqu.On(goqu.Ex{"comment.id_task": goqu.I("task.id")})).Where(goqu.I("service_name.id").In(m.Id)).ScanVals(&commentFile)
	if err != nil {
		tx.Rollback()
		return FileName{}, err
	}

	// Выгрузка список файлов закрепленными за сервисами
	var serviceFile = make([]string, 0)
	err = tx.Select("service.file_name").From("service").Join(goqu.T("service_name"), goqu.On(goqu.Ex{"service_name.id": goqu.I("service.id_name")})).Where(goqu.I("service_name.id").In(m.Id)).ScanVals(&serviceFile)
	if err != nil {
		tx.Rollback()
		return FileName{}, err
	}

	if err := tx.Commit(); err != nil {
		return FileName{}, err
	}

	return FileName{Comment: commentFile, Service: serviceFile}, nil
}
