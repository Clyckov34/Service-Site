package page_manager

import (
	"database/sql"
	"net/http"
	"repair/pkg/check"
	"repair/pkg/errs"
	"repair/pkg/session"
	"repair/service/admin/page_manager"
	"repair/service/admin/page_setting"
	"repair/service/install/page_reg"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page главная страница
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		var manager = page_manager.Param{Id: sess.Id.(int)}
		data, err := manager.GetManager()
		if err != nil {
			data = nil
		}

		ctx.HTML(http.StatusOK, "cabinet_manager.html", gin.H{
			"Manager": data,
		})

	})
}

// Detail подробно о менеджере
func (m *Page) Detail(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if sess.Id.(int) == id {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

		var setting = page_setting.Param{Id: id}

		res, err := setting.GetSetting()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_manager_detail.html", gin.H{
			"Id":       id,
			"Login":    res.Login,
			"FullName": res.FullName,
			"Email":    res.Email,
		})

	})
}

// Add добавляет менеджера
func (m *Page) Add(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		tp, err := strconv.ParseBool(ctx.PostForm("Type"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var manager = page_reg.Manager{
			Login:     strings.TrimSpace(ctx.PostForm("Login")),
			Password:  strings.TrimSpace(ctx.PostForm("Password")),
			FirstName: strings.TrimSpace(ctx.PostForm("FirstName")),
			Email:     strings.TrimSpace(ctx.PostForm("Email")),
			Admin:     tp,
		}

		if err := manager.RegManager(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if err := manager.MailSend(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/manager")
	})
}

// EditPassword рактирует пароль
func (m *Page) EditPassword(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		newPass := strings.TrimSpace(ctx.PostForm("NewPassword"))
		repeat := strings.TrimSpace(ctx.PostForm("Repeat"))

		//Проверяет новый и повторяющий пароль
		if newPass != repeat {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, "Пароль не совпадает")
			return
		}

		id, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if sess.Id.(int) == id {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

		var setting = &page_setting.Param{Id: id}

		if err := setting.EditPassword(newPass); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/manager")

	})
}

// EditEmail редактирует логин
func (m *Page) EditEmail(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		id, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if sess.Id.(int) == id {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

		var setting = page_setting.Param{
			Id:    id,
			Email: strings.TrimSpace(ctx.PostForm("Email")),
		}

		if err := setting.EditEmail(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/manager")
	})
}

// EditLogin редактирует логин
func (m *Page) EditLogin(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		id, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if sess.Id.(int) == id {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

		var setting = page_setting.Param{
			Id:    id,
			Login: strings.TrimSpace(ctx.PostForm("Login")),
		}

		if err := setting.EditLogin(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/manager")

	})
}

// EditFullName редактирует ФИО
func (m *Page) EditFullName(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		id, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if sess.Id.(int) == id {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

		var setting = page_setting.Param{
			Id:       id,
			FullName: strings.TrimSpace(ctx.PostForm("FullName")),
		}

		if err := setting.EditFullName(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/manager")

	})
}

// EditType редактирует тип учетную записб
func (m *Page) EditType(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		id, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		tp, err := strconv.ParseBool(ctx.PostForm("Type"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var manager = page_manager.Param{
			Id:   id,
			Type: sql.NullBool{Bool: tp, Valid: true},
		}

		if err := manager.EditType(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/manager")
	})
}

// Delete удаляет учетную запись
func (m *Page) Delete(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		id, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if sess.Id.(int) == id {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

		var setting = page_manager.Param{Id: id}

		if err := setting.Delete(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/manager")

	})
}
