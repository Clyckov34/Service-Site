package page_form

import (
	"net/http"
	"repair/pkg/check"
	"repair/pkg/errs"
	"repair/service/install/page_reg"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page форма для заполнения полей перед установкой
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		id, err := check.InstallTable("auth")
		if err == nil || id != 0 {
			ctx.Redirect(http.StatusMovedPermanently, "/admin")
			return
		}

		ctx.HTML(http.StatusOK, "install_form.html", nil)

	})
}

// Register регистрация
func (m *Page) Register(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		var create = &page_reg.Manager{
			Login:     strings.TrimSpace(ctx.PostForm("Login")),
			Password:  strings.TrimSpace(ctx.PostForm("Password")),
			FirstName: strings.TrimSpace(ctx.PostForm("FirstName")),
			Email:     strings.TrimSpace(ctx.PostForm("Email")),
			Admin:     true,
		}

		if err := create.CreateTable(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if err := create.RegStatusList(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if err := create.RegIcon(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if err := create.RegManager(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if err := create.MailSend(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin")

	})
}
