package page_back_mail

import (
	"net/http"
	"repair/pkg/errs"
	"repair/service/lending/page_back_mail"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Send запрос на обратную связь. Отправка на почту
func (m *Page) Send(url string) {
	m.Router.POST(url, func(ctx *gin.Context) {

		var mail = page_back_mail.Param{
			FirstName: strings.TrimSpace(ctx.PostForm("FirstName")),
			Phone:     strings.TrimSpace(ctx.PostForm("Phone")),
			Email:     strings.TrimSpace(ctx.PostForm("Email")),
			Text:      strings.TrimSpace(ctx.PostForm("Text")),
		}

		if err := mail.Send(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/")

	})

}
