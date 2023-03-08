package page_filter

import (
	"net/http"
	"repair/pkg/errs"
	"repair/pkg/session"
	"repair/service/admin/page_filter"
	"repair/service/admin/page_task"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page страница главная
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		statusChan := make(chan page_filter.StatusChan)
		categoryChan := make(chan page_filter.CategoryChan)
		managerChan := make(chan page_task.ManagerChan)

		go page_filter.GetStatus(statusChan)
		go page_filter.GetCategory(categoryChan)

		var task = page_task.Param{}
		go task.GetManager(managerChan)

		status := <-statusChan
		category := <-categoryChan
		manager := <-managerChan

		if status.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if category.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if manager.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_task_filter.html", gin.H{
			"StatusList":   status.Data,
			"CategoryList": category.Data,
			"ManagerList":  manager.Data,
			"DateStart":    time.Now().Format("2006-01") + "-01",
			"DateEnd":      time.Now().Format("2006-01-02"),
		})

	})
}

// Search поиск задач
func (m *Page) Search(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		task, err := page_filter.ParserNumberTask(strings.TrimSpace(ctx.DefaultQuery("Task", "")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var filter = page_filter.Filter{
			Status:    ctx.DefaultQuery("Status", ""),
			Category:  ctx.DefaultQuery("Category", ""),
			FirstName: strings.TrimSpace(ctx.DefaultQuery("Family", "")),
			Phone:     strings.TrimSpace(ctx.DefaultQuery("Phone", "")),
			Task:      task,
			Address:   strings.TrimSpace(strings.TrimSpace(ctx.DefaultQuery("Address", ""))),
			Manager:   ctx.DefaultQuery("Manager", ""),
			DateStart: ctx.DefaultQuery("DateStart", ""),
			DateEnd:   ctx.DefaultQuery("DateEnd", ""),
		}

		data, err := filter.GetTaskFilter()
		if err != nil {
			data = nil
		}

		ctx.JSON(http.StatusOK, data)

	})

}
