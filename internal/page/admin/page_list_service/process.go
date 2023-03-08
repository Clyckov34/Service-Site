package page_list_service

import (
	"net/http"
	"repair/pkg/check"
	"repair/pkg/errs"
	fl "repair/pkg/file"
	"repair/pkg/session"
	"repair/service/admin/page_list_service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

const Dir = "static/images/service/"

// Page главная страница список услуг
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

		typeRepair := make(chan page_list_service.TypeRepairData)
		priceListService := make(chan page_list_service.ServicePriceListData)

		go page_list_service.GetTypeRepair(typeRepair)
		go page_list_service.GetPriceListService(priceListService)

		tpRepair := <-typeRepair
		plService := <-priceListService

		if tpRepair.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		if plService.Error != nil {
			plService.Data = nil
		}

		ctx.HTML(http.StatusOK, "cabinet_list_service.html", gin.H{
			"TypeRepair": tpRepair.Data,
			"Services":   plService.Data,
		})

	})

}

// Detail подробная услуга
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

		id, err := strconv.Atoi(strings.TrimSpace(ctx.Param("id")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		var service = &page_list_service.Param{Id: id}

		res, err := service.GetServiceId()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_list_service_detail.html", gin.H{
			"Title":    res.Title,
			"Price":    res.Price,
			"Sale":     res.Sale,
			"FileName": res.FileName,
			"Text":     res.Text,
			"Id":       res.Id,
		})

	})

}

// Add добавляет услугу
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

		typeRepair, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("IdName")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		price, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Price")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		file, err := ctx.FormFile("File")
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		nameFile, err := fl.UploadAndCompress(ctx, file, Dir)
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		sale, _ := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Sale")))

		var service = &page_list_service.Param{
			Type:     typeRepair,
			Title:    strings.TrimSpace(ctx.PostForm("Title")),
			Text:     strings.TrimSpace(ctx.PostForm("Text")),
			FileName: nameFile,
			Price:    price,
			Sale:     sale,
		}

		if err := service.Add(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/list-service")

	})

}

// Edit изменяет данные
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
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		price, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Price")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		sale, _ := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Sale")))

		var service = &page_list_service.Param{
			Id:    id,
			Title: strings.TrimSpace(ctx.PostForm("Title")),
			Text:  strings.TrimSpace(ctx.PostForm("Text")),
			Price: price,
			Sale:  sale,
		}

		if err := service.Edit(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/list-service")

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

		var service = page_list_service.Param{
			Id:       id,
			FileName: strings.TrimSpace(ctx.PostForm("FileName")),
		}

		// Удаляет файл
		fl.Delete(service.FileName, Dir)

		// Загрузка файла
		nameFile, err := fl.UploadAndCompress(ctx, file, Dir)
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		// Изменяет данные в бд
		if err := service.UpdatePortfolio(nameFile); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/list-service")

	})
}

// Delete удаляет запись о услуге
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
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		var service = &page_list_service.Param{
			Id:       id,
			FileName: strings.TrimSpace(ctx.PostForm("FileName")),
		}

		if err := service.Delete(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		fl.Delete(service.FileName, Dir)

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/list-service")

	})

}

// Search поиск
func (m *Page) Search(url string) {
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

		search := ctx.DefaultQuery("p", "")

		var service = &page_list_service.Param{Title: search}

		res, err := service.Search()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, res)

	})

}
