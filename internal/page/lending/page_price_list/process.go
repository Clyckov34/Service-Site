package page_price_list

import (
	"net/http"
	"repair/service/lending/page_price_list"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

//PriceList список услуг
func (m *Page) PriceList(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		res, err := page_price_list.GetServiceAll()
		if err != nil {
			res = nil
		}

		ctx.HTML(http.StatusOK, "price_list.html", gin.H{
			"Title":    "HOME - Услуги для дома | Услуги",
			"Services": res,
		})

	})
}

//PriceListSearch раздел услуги | Поиск JSON
func (m *Page) PriceListSearch(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		search := ctx.DefaultQuery("p", "")

		res, err := page_price_list.GetServiceName(search)
		if err != nil {
			res = nil
		}

		ctx.JSON(http.StatusOK, res)
	})

}
