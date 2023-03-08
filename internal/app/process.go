package app

import (
	"repair/router"
)

//Run запуск приложения
func Run() error {
	setting := server()

	static(setting.Engine)
	router.List(setting.Engine)

	if err := setting.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
