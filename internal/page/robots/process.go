package robots

import "github.com/gin-gonic/gin"

type Page struct {
	Router *gin.RouterGroup
}

//RobotsTXT описание робота
func (m *Page) RobotsTXT(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		ctx.File("robots.txt")

	})
}
