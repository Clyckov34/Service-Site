package page_portfolio

import (
	"net/http"
	"repair/pkg/check"
	"repair/pkg/errs"
	fl "repair/pkg/file"
	"repair/pkg/session"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

const dir = "static/images/portfolio/"

// Page Портфолио форма
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

		res, err := fl.GetFileAll(dir)
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_portfolio.html", gin.H{"Data": res})

	})

}

// Add добавляет портфолио
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

		form, _ := ctx.MultipartForm()
		files := form.File["Files"]

		err = fl.UploadAllAndCompress(ctx, files, dir)
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/portfolio")

	})
}

// Delete Удаляет фото
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

		file := ctx.PostForm("Delete")

		if err := fl.Delete(file, dir); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/portfolio")

	})

}
