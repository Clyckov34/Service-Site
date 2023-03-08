package page_cabinet

import (
	"net/http"
	"repair/pkg/session"
	"repair/service/admin/page_cabinet"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page главная страница личного кабинета
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		check, err := page_cabinet.CheckTask("To Do")
		if err != nil {
			check = false
		}

		ctx.HTML(http.StatusOK, "cabinet_home.html", gin.H{
			"Task":  check,
			"Admin": sess.Admin,
		})

	})

}
