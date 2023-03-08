package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Setting struct {
	Engine *gin.Engine
	Server *http.Server
}

// server настройка сервера
func server() Setting {
	engine := gin.Default()
	engine.MaxMultipartMemory = 50 << 20

	server := &http.Server{
		Addr:           ":8888",
		Handler:        engine,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return Setting{Engine: engine, Server: server}
}

// static Статические файлы
func static(r *gin.Engine) {
	r.LoadHTMLGlob("static/html/*")
	// r.Static("/img/", "./static/images")
	// r.Static("/js/", "./static/javascript")
	// r.Static("/css/", "./static/css")
}
