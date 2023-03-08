package page_setting

import (
	"net/http"
	"repair/pkg/errs"
	"repair/pkg/session"
	"repair/service/admin/page_setting"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page страница настроек админа
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		var setting = &page_setting.Param{Id: sess.Id.(int)}

		res, err := setting.GetSetting()
		if err != nil {
			res = page_setting.Param{}
		}

		ctx.HTML(http.StatusOK, "cabinet_setting.html", gin.H{
			"Login":    res.Login,
			"FullName": res.FullName,
			"Email":    res.Email,
		})

	})

}

// EditEmail Изменяет email
func (m *Page) EditEmail(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		var setting = &page_setting.Param{
			Email: strings.TrimSpace(ctx.PostForm("Email")),
			Id:    sess.Id.(int),
		}

		if err := setting.EditEmail(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(301, "/admin/exit")

	})

}

// EditFullName Изменяет ФИО
func (m *Page) EditFullName(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		var setting = &page_setting.Param{
			FullName: strings.TrimSpace(ctx.PostForm("FullName")),
			Id:       sess.Id.(int),
		}

		if err := setting.EditFullName(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(301, "/admin/cabinet")

	})

}

// EditLogin изменяет логин
func (m *Page) EditLogin(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		login := strings.TrimSpace(ctx.PostForm("Login"))

		var setting = &page_setting.Param{
			Login: login,
			Id:    sess.Id.(int),
		}

		if err := setting.EditLogin(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(301, "/admin/exit")

	})
}

// EditPassword изменяет раполь
func (m *Page) EditPassword(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		nowPass := strings.TrimSpace(ctx.PostForm("NowPassword"))
		newPass := strings.TrimSpace(ctx.PostForm("NewPassword"))
		repeat := strings.TrimSpace(ctx.PostForm("Repeat"))

		//Проверяет новый и повторяющий пароль
		if newPass != repeat {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, "Пароль не совпадает")
			return
		}

		var setting = &page_setting.Param{Id: sess.Id.(int)}

		//Изменяет пароль
		if err := setting.EditCheckPassword(nowPass, newPass); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(301, "/admin/exit")
	})

}
