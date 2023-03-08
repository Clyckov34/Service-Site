package page_static

import (
	"net/http"
	"repair/pkg/session"
	"repair/pkg/errs"
	"repair/service/admin/page_static"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page главная страница
func (m *Page) Page(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		_, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		res, err := page_static.GetStaticsAll()
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_static.html", gin.H{
			"AllToDo":         res.ToDo,
			"AllInProgress":   res.InProgress,
			"AllPause":        res.Pause,
			"AllDenied":       res.Denied,
			"AllDone":         res.Done,
		})

	})
}
