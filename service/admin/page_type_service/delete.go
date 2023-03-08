package page_type_service

import (
	"repair/pkg/database"
	fl "repair/pkg/file"
	"sync"

	"github.com/doug-martin/goqu/v9"
)

// DeletePortfolio удаляет профиль
func (m *Param) DeletePortfolio() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Delete("service_name").Where(goqu.I("id").In(m.Id)).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}


// DeleteFileProfile удаляет файлы от привязанных сервисов
func (m *Param) DeleteFileProfile(dirComment, dirService string) error {
	fileList, err := m.getProfileFile()
	if err != nil {
		return err
	}

	var file = WG{sync.WaitGroup{}}
	file.wp.Add(2)
		
	go file.deleteFile(fileList.Comment, dirComment)
	go file.deleteFile(fileList.Service, dirService)

	file.wp.Wait()
	return nil
}

// deleteFile удаляет файлы
func (m *WG) deleteFile(files []string, dir string) {
	defer m.wp.Done() 

	for _, f := range files {
		if err := fl.Delete(f, dir); err != nil {
			continue
		}
	}
}
