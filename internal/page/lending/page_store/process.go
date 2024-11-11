package page_store

import (
	"net/http"
	"repair/pkg/check"
	"repair/pkg/errs"
	"repair/service/lending/page_store"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page форма заказа услуги
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		category, err := check.Сategory(ctx.Param("name"))
		if err != nil {
			ctx.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		data, err := page_store.GetServiceName(category.Id)
		if err != nil {
			data = nil
		}

		ctx.HTML(http.StatusOK, "store.html", gin.H{
			"Title":          "HOME - Услуги для дома",
			"Type_Repair":    category.Title,
			"Id_Type_Repair": category.Id,
			"Service":        data,
		})

	})
}

// Detail подробно об услуге
func (m *Page) Detail(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		idService, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			ctx.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		store, err := page_store.GetServiceID(ctx.Param("name"), idService)
		if err != nil {
			ctx.Redirect(http.StatusMovedPermanently, "/")
			return
		}

		ctx.HTML(http.StatusOK, "store_detail.html", gin.H{
			"Category":    store.CategoryTitle,
			"Category_ID": store.CategoryID,
			"Service":     store.ServiceTitle,
			"Service_ID":  store.ServiceId,
			"Price":       store.ServicePrice,
			"Sale":        store.ServiceSale,
			"Text":        store.ServiceText,
			"Img":         store.ServiceFileName,
		})

	})

}

// Send отправка заявки на ремонт на mail и в бд
func (m *Page) Send(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {
		IdCategory, err := strconv.Atoi(ctx.PostForm("IdCategory"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		idService, err := strconv.Atoi(ctx.PostForm("IdService"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var task = page_store.Send{
			IdCategory: IdCategory,
			IdService:  idService,
			Street:     strings.TrimSpace(ctx.PostForm("Street")),
			FirstName:  strings.TrimSpace(ctx.PostForm("FirstName")),
			Phone:      strings.TrimSpace(ctx.PostForm("Phone")),
			Email:      strings.TrimSpace(ctx.PostForm("Email")),
			Date:       time.Now().Format("2006-01-02 15:04:05"),
		}

		store, err := task.CreateTask("To Do")
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if err := task.SendMail(store); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/")

	})

}
