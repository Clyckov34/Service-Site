package page_type_service

import (
	"net/http"
	"repair/internal/page/admin/page_comment"
	"repair/internal/page/admin/page_list_service"
	"repair/pkg/check"
	"repair/pkg/errs"
	fl "repair/pkg/file"
	"repair/pkg/session"
	"repair/service/admin/page_type_service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

const dir = "static/images/type_repair/"

// Page Виды услуг
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

		var profile = page_type_service.Param{}
		res, err := profile.GetProfile()
		if err != nil {
			res = nil
		}

		ctx.HTML(http.StatusOK, "cabinet_type_service.html", gin.H{"Data": res})

	})
}

// Add Добавляет профиль
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

		file, err := ctx.FormFile("File")
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		var profile = &page_type_service.Param{
			Title:   strings.TrimSpace(ctx.PostForm("Title")),
			Url:     strings.TrimSpace(ctx.PostForm("URL")),
			KeyType: strings.TrimSpace(ctx.PostForm("Key")),
		}

		nameFile, err := fl.UploadAndCompress(ctx, file, dir)
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if err = profile.CreateService(nameFile); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/type-service")

	})
}

// Edit редактирует данные
func (m *Page) Edit(url string) {
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

		id, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Id")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		var profile = &page_type_service.Param{
			Id:      id,
			Title:   strings.TrimSpace(ctx.PostForm("Title")),
			Url:     strings.TrimSpace(ctx.PostForm("URL")),
			KeyType: strings.TrimSpace(ctx.PostForm("Key")),
		}

		if err := profile.Edit(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/type-service")
	})
}

// EditPhoto редактирует фото
func (m *Page) EditPhoto(url string) {
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

		id, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Id")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		file, err := ctx.FormFile("File")
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		var profile = &page_type_service.Param{
			Id:       id,
			FileName: strings.TrimSpace(ctx.PostForm("FileName")),
		}

		// Удаляет файл
		fl.Delete(profile.FileName, dir)

		// Загрузка файла
		nameFile, err := fl.UploadAndCompress(ctx, file, dir)
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		// Изменяет данные в бд
		if err := profile.UpdatePortfolio(nameFile); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/type-service")
	})
}

// Delete удаляет услугу
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

		id, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Id")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		var profile = &page_type_service.Param{
			Id:       id,
			FileName: strings.TrimSpace(ctx.PostForm("FileName")),
		}

		if err := profile.DeletePortfolio(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		fl.Delete(profile.FileName, dir)                                   // Удаляет фото портфолио
		profile.DeleteFileProfile(page_comment.Dir, page_list_service.Dir) // Удаляет фото привязаных к портфолио

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/type-service")
	})
}

// Detail подробное
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
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		var profile = &page_type_service.Param{Id: id}

		res, err := profile.GetProfileId()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_type_service_detail.html", gin.H{
			"Id":       res.Id,
			"Title":    res.Title,
			"Url":      res.Url,
			"FileName": res.FileName,
			"KeyType":  res.KeyType,
		})
	})
}
