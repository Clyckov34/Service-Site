package history

import (
	"log"
	"repair/pkg/database"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type New struct {
	IdManager int
	IdTask    int
	IP        string
}

//Write запись истории
func (m *New) Write(title string) {
	db, err := database.Open()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	dialect := database.Dialect(db)
	_, err = dialect.Insert("history_task").Rows(goqu.Record{
		"id_auth": m.IdManager,
		"id_task": m.IdTask,
		"title":   title,
		"date":    nowDateTime(),
		"ip":      m.IP,
	}).Executor().Exec()
	if err != nil {
		log.Println(err)
	}
}

// nowDateTime выводит текущие дате и время
func nowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
