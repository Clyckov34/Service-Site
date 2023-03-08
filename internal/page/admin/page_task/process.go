package page_task

import (
	"fmt"
	"net/http"
	pc "repair/internal/page/admin/page_comment"
	"repair/pkg/check"
	"repair/pkg/errs"
	"repair/pkg/history"
	"repair/pkg/session"
	"repair/service/admin/page_comment"
	"repair/service/admin/page_task"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page главная страница заявки
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		taskToDo := make(chan page_task.TaskChan)
		taskInProgress := make(chan page_task.TaskChan)
		taskPause := make(chan page_task.TaskChan)
		taskDenied := make(chan page_task.TaskChan)
		taskDone := make(chan page_task.TaskChan)

		var limit uint = 50 // лимит на выгрузку кол-во задач

		go page_task.GetTaskName("To Do", limit, taskToDo)
		go page_task.GetTaskName("In Progress", limit, taskInProgress)
		go page_task.GetTaskName("Pause", limit, taskPause)
		go page_task.GetTaskName("Denied", limit, taskDenied)
		go page_task.GetTaskName("Done", limit, taskDone)

		todo := <-taskToDo
		inProgress := <-taskInProgress
		pause := <-taskPause
		denied := <-taskDenied
		done := <-taskDone

		if todo.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if inProgress.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if pause.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if denied.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if done.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_task.html", gin.H{
			"ToDo":       todo.Data,
			"InProgress": inProgress.Data,
			"Pause":      pause.Data,
			"Denied":     denied.Data,
			"Done":       done.Data,
		})

	})
}

// Detail подробно о задачке
func (m *Page) Detail(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var task = page_task.Param{Id: id}

		statusChan := make(chan page_task.StatusChan)
		taskNameChan := make(chan page_task.TaskNameChan)
		commentChan := make(chan page_task.CommentChan)
		managerChan := make(chan page_task.ManagerChan)
		taskHistoryChan := make(chan page_task.TaskHistoryChan)

		go page_task.GetStatus(statusChan)
		go task.GetTaskNameDetail(taskNameChan)
		go task.GetComment(commentChan)
		go task.GetManager(managerChan)
		go task.GetTaskHistory(taskHistoryChan)

		status := <-statusChan
		tasks := <-taskNameChan
		comment := <-commentChan
		manager := <-managerChan
		taskHist := <-taskHistoryChan

		if status.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, status.Error.Error())
			return
		}

		if tasks.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, tasks.Error.Error())
			return
		}

		if comment.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, comment.Error.Error())
			return
		}

		if manager.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, manager.Error.Error())
			return
		}

		if taskHist.Error != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, manager.Error.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_task_detail.html", gin.H{
			"Category":         tasks.Data.Category,
			"KeyType":          tasks.Data.KeyType,
			"Title":            tasks.Data.Title,
			"FirstName":        tasks.Data.FirstName,
			"FirstNameManager": tasks.Data.FirstNameManager,
			"Phone":            tasks.Data.Phone,
			"Email":            tasks.Data.Email,
			"Address":          tasks.Data.Address,
			"Price":            tasks.Data.Price,
			"Sale":             tasks.Data.Sale,
			"PriceWork":        tasks.Data.PriceWork,
			"Status":           tasks.Data.Status,
			"StatusTranslate":  tasks.Data.StatusTranslate,
			"DateStart":        tasks.Data.DateStart,
			"DateStatus":       tasks.Data.DateStatus,
			"IdTask":           id,
			"FileName":         tasks.Data.FileName,
			"StatusList":       status.Data,
			"CommentList":      comment.Data,
			"ManagerList":      manager.Data,
			"TaskHistory":      taskHist.Data,
		})

	})

}

// DetailEdit подробно о редактирование задачке
func (m *Page) DetailEdit(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var task = page_task.Param{Id: id}

		data, err := task.GetTaskNameDetailId()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_task_edit.html", gin.H{
			"Id":        id,
			"KeyType":   data.KeyType,
			"FirstName": data.FirstName,
			"Phone":     data.Phone,
			"Email":     data.Email,
			"Address":   data.Address,
			"Price":     data.Price,
		})

	})

}

// EditStatus Изменяет статус
func (m *Page) EditStatus(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		id, err := strconv.Atoi(ctx.PostForm("Status"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var status = page_task.Param{
			Id:       id,
			IdTask:   idTask,
			DateTime: time.Now().Format("2006-01-02 15:04:05"),
		}

		nameStatus, err := status.EditStatus()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		taskHist := history.New{
			IdManager: sess.Id.(int),
			IdTask:    idTask,
			IP:        sess.Ip.(string),
		}

		go taskHist.Write(fmt.Sprintf(`Изменен статус: %v`, nameStatus))

		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", status.IdTask))

	})

}

// EditPrice редактирует и добавляет цену
func (m *Page) EditPrice(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		price, err := strconv.Atoi(ctx.PostForm("Price"))
		if err != nil {
			price = 0
		}

		var task = page_task.Param{
			Id:    idTask,
			Price: price,
		}

		if err := task.EditPrice(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		taskHist := history.New{
			IdManager: sess.Id.(int),
			IdTask:    idTask,
			IP:        sess.Ip.(string),
		}

		go taskHist.Write(fmt.Sprintf("Добавлен (Изменен) стоимость работ: %v руб", price))

		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", task.Id))

	})

}

// EditTask редактирует задачу
func (m *Page) EditTask(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		idTask, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		price, err := strconv.Atoi(ctx.PostForm("Price"))
		if err != nil {
			price = 0
		}

		var task = page_task.Param{
			Id:        idTask,
			FirstName: ctx.PostForm("FirstName"),
			Phone:     ctx.PostForm("Phone"),
			Email:     ctx.PostForm("Email"),
			Address:   ctx.PostForm("Address"),
			Price:     price,
		}

		if err := task.EditTask(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		taskHist := history.New{
			IdManager: sess.Id.(int),
			IdTask:    idTask,
			IP:        sess.Ip.(string),
		}

		go taskHist.Write(fmt.Sprintf("Изменена задача: ФИО Заказчика: %v. Телефон: %v. Почта: %v. Адрес: %v. Цена: %v", task.FirstName, task.Phone, task.Email, task.Address, task.Price))

		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", task.Id))

	})

}

// EditManager редактирует добавляет менеждера в задачку
func (m *Page) EditManager(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		idManager, err := strconv.Atoi(ctx.PostForm("Manager"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		var task = page_task.Param{
			IdTask:             idTask,
			FirstNameManagerId: idManager,
		}

		managerName, err := task.EditManager()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		taskHist := history.New{
			IdManager: sess.Id.(int),
			IdTask:    idTask,
			IP:        sess.Ip.(string),
		}

		go taskHist.Write(fmt.Sprintf("Добавлен (Изменен) исполнитель: %v", managerName))

		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", task.IdTask))
	})
}

// Delete удаляет задачу
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

		var task = page_task.Param{Id: id}

		fileList, _ := page_comment.GetCommentPhoto(task.Id)
		_ = task.DeleteFile(fileList, pc.Dir)

		if err := task.DeleteTask(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/task")

	})

}
