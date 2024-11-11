package page_home

import (
	"net/http"
	"repair/service/lending/page_home"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Home Главная страница
func (m *Page) Home(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		service := make(chan page_home.ServiceData)
		socialNetwork := make(chan page_home.SocialNetworkData)
		portfolio := make(chan page_home.PortfolioData)

		go page_home.GetService(service)
		go page_home.GetSocialNetwork(socialNetwork)
		go page_home.GetPortfolio(portfolio)

		ser := <-service
		socNet := <-socialNetwork
		portf := <-portfolio

		if ser.Error != nil {
			ser.Data = nil
		}

		if socNet.Error != nil {
			socNet.Data = nil
		}

		if portf.Error != nil {
			portf.Data = nil
		}

		ctx.HTML(http.StatusOK, "home.html", gin.H{
			"Title":         "HOME - Услуги для дома в Суровикино",
			"MapsUrl":       "https://yandex.ru/maps/10980/surovikino/?from=mapframe&ll=42.842342%2C48.602958&mode=usermaps&source=mapframe&um=constructor%3Afc0342683d01c36e96461298b9c301d882d4e8324aed120ee63ed702fd121395&utm_source=mapframe&z=17",
			"Service":       ser.Data,
			"Portfolio":     portf.Data,
			"SocialNetwork": socNet.Data,
		})

	})

}

// Review отзывы
func (m *Page) Review(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "review.html", gin.H{
			"Title": "HOME - Услуги для дома | Отзывы",
		})

	})

}
