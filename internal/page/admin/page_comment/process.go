package page_comment

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"net/http"
	"repair/pkg/check"
	"repair/pkg/errs"
	fl "repair/pkg/file"
	"repair/pkg/session"
	"repair/service/admin/page_comment"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

const Dir = "static/images/comment/"

func (m *Page) Detail(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		idMngComm, err := page_comment.GetIdManagerComment(id)
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		//Доступ имеет Админ и комментирующий
		if sess.Id == idMngComm || check.ManagerAdmin(sess.Admin.(bool)) == nil {
			data, err := page_comment.GetCommentId(id)
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			//Установка сессии
			s := sessions.Default(ctx)
			s.Set("IdMngComm", idMngComm)
			if err = s.Save(); err != nil {
				errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			ctx.HTML(http.StatusOK, "cabinet_comment_detail.html", gin.H{
				"FileName": data.FileName,
				"Id":       data.Id,
				"IdTask":   data.IdTask,
				"Text":     data.Text,
			})
		} else {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

	})
}

// Add Добавляет комментарии к задачке
func (m *Page) Add(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		file, err := ctx.FormFile("File")
		if err != nil {
			file = nil
		}

		var comment = page_comment.Comment{
			Id:        idTask,
			IdManager: sess.Id.(int),
			Text:      ctx.PostForm("Comment"),
			Date:      time.Now().Format("2006-01-02 15:04:05"),
		}

		var nameFile string

		if file != nil {
			name, err := fl.UploadAndCompress(ctx, file, Dir)
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			nameFile = name
		}

		if err := comment.AddBD(nameFile); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", idTask))

	})
}

// Edit редактирует комментарий
func (m *Page) Edit(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if sess.Id == sess.IdMngComm || check.ManagerAdmin(sess.Admin.(bool)) == nil {
			id, err := strconv.Atoi(ctx.PostForm("Id"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			var comment = page_comment.Message{
				Id:   id,
				Text: ctx.PostForm("Text"),
			}

			if err := comment.Edit(); err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", idTask))
		} else {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

	})
}

// EditPhoto меняет фото удаляя старое в замен на новое
func (m *Page) EditPhoto(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if sess.Id == sess.IdMngComm || check.ManagerAdmin(sess.Admin.(bool)) == nil {
			id, err := strconv.Atoi(ctx.PostForm("Id"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			file, err := ctx.FormFile("File")
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			var comment = page_comment.Comment{Id: id}

			// Удаляет файл
			fl.Delete(ctx.PostForm("FileName"), Dir)

			// Загружает файл
			namefile, err := fl.UploadAndCompress(ctx, file, Dir)
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			if err := comment.EditNameFileDB(namefile); err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", idTask))
		} else {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

	})
}

// DeletePhoto удалить фото
func (m *Page) DeletePhoto(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if sess.Id == sess.IdMngComm || check.ManagerAdmin(sess.Admin.(bool)) == nil {
			id, err := strconv.Atoi(ctx.PostForm("Id"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			var comment = page_comment.Message{
				Id:     id,
				IdTask: idTask,
			}

			fl.Delete(ctx.PostForm("FileName"), Dir)

			if err := comment.DeleteNameFileDB(); err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", comment.IdTask))
		} else {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

	})
}

// Delete удаляет комментарий
func (m *Page) Delete(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		if sess.Id == sess.IdMngComm || check.ManagerAdmin(sess.Admin.(bool)) == nil {
			id, err := strconv.Atoi(ctx.PostForm("Id"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			idTask, err := strconv.Atoi(ctx.PostForm("IdTask"))
			if err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			var comment = page_comment.Message{
				Id:     id,
				IdTask: idTask,
			}

			// Удаляет файлы фото если Они есть
			_ = fl.Delete(ctx.PostForm("FileName"), Dir)

			if err := comment.DeleteComment(); err != nil {
				errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
				return
			}

			//Удаляет сессию
			s := sessions.Default(ctx)
			s.Delete("IdMngComm")
			if err = s.Save(); err != nil {
				errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			ctx.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/admin/cabinet/task/%v", comment.IdTask))
		} else {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, "доступ запрещен")
			return
		}

	})
}
