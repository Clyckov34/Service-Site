package session

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

type SessGet struct {
	Ip        interface{}
	Id        interface{}
	Email     interface{}
	Admin     interface{}
	IdMngComm interface{}
}

// Check проверяет  сессии на наличии параметров
func Check(ctx *gin.Context) (SessGet, error) {
	sess := sessions.Default(ctx)

	ip := sess.Get("Ip")
	id := sess.Get("Id")
	email := sess.Get("Email")
	admin := sess.Get("Admin")
	idMngComm := sess.Get("IdMngComm")

	if id == nil || ip != ctx.ClientIP() || email != true {
		return SessGet{Ip: ip, Id: id, Email: email, Admin: admin, IdMngComm: idMngComm}, errors.New("параметры в сессиях не совпадает")
	}

	return SessGet{Ip: ip, Id: id, Email: email, Admin: admin, IdMngComm: idMngComm}, nil
}

// Options параметры сессии
func Options(maxAge int) sessions.Options {
	return sessions.Options{
		Path:     "/",
		Domain:   os.Getenv("DOMAIN"),
		MaxAge:   maxAge, 
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

// StoreKeyGenerate генерация ключа шифрование session
func StoreKeyGenerate() cookie.Store {
	store := cookie.NewStore(securecookie.GenerateRandomKey(32))
	store.Options(Options(86400)) //Сессия на сутки
	return store
}
