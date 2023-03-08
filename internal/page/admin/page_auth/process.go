package page_auth

import (
	"net/http"
	"repair/pkg/errs"
	ml "repair/pkg/mail"
	"repair/pkg/session"
	"repair/service/admin/page_auth"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page главная страница
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err == nil {
			ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet")
		}

		ctx.HTML(http.StatusOK, "auth_form.html", nil)

	})
}

// Auth авторизация менееджера
func (m *Page) Auth(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err == nil {
			ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet")
			return
		}

		var manager = page_auth.Page{
			Login:    strings.TrimSpace(ctx.PostForm("Login")),
			Password: strings.TrimSpace(ctx.PostForm("Password")),
		}

		//Авторизация
		user, err := manager.AuthCabinet()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		//Отправляет письмо
		if err := ml.Send(user.Manager.Email, "Подтверждающий код", page_auth.Message(user.Security.Code, user.Security.DateTime), nil); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		//Установка сессии
		s := sessions.Default(ctx)
		s.Set("Id", user.Manager.Id)
		s.Set("Ip", ctx.ClientIP())
		s.Set("Admin", user.Manager.Admin)
		if err = s.Save(); err != nil {
			ctx.Redirect(302, "/admin")
			return
		}

		ctx.HTML(http.StatusOK, "auth_check_email.html", gin.H{"Email": user.Manager.Email})

	})
}

// CheckToMail проверка через почту
func (m *Page) CheckToMail(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err == nil {
			ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet")
			return
		}

		check := page_auth.Check{
			Id:       sess.Id.(int),
			Code:     strings.TrimSpace(ctx.PostForm("Code")),
			DateTime: time.Now().Format("2006-01-02 15:04:05"),
		}

		//Проверяем проверяющие данные на актуальность
		if err := check.CodeAndDateTime(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		//Установка сессии
		s := sessions.Default(ctx)
		s.Set("Email", true)
		if err := s.Save(); err != nil {
			ctx.Redirect(302, "/admin")
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet")

	})
}

// SessionExit выход из сессии
func (m *Page) SessionExit(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		s := sessions.Default(ctx)
		s.Clear()
		s.Options(session.Options(-1))
		s.Save()

		ctx.Redirect(http.StatusTemporaryRedirect, "/admin")

	})
}
