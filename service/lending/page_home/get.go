package page_home

import (
	"repair/pkg/database"
	"repair/pkg/file"

	"github.com/doug-martin/goqu/v9"
)

// GetService получить сервисы
func GetService(ch chan ServiceData) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- ServiceData{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var service = make([]Service, 0)
	if err = dialect.Select(&service).From("service_name").ScanStructs(&service); err != nil {
		ch <- ServiceData{nil, err}
		return
	}

	ch <- ServiceData{service, nil}
}

// GetSocialNetwork выгрузка групп социальных сетей
func GetSocialNetwork(ch chan SocialNetworkData) {
	defer close(ch)

	db, err := database.Open()
	if err != nil {
		ch <- SocialNetworkData{nil, err}
		return
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var socialNetwork = make([]SocialNetwork, 0)
	if err = dialect.Select(&socialNetwork).From("social_network").Join(goqu.T("icon"), goqu.On(goqu.Ex{"social_network.id_icon": goqu.I("icon.id")})).ScanStructs(&socialNetwork); err != nil {
		ch <- SocialNetworkData{nil, err}
		return
	}

	ch <- SocialNetworkData{socialNetwork, nil}
}

// GetPortfolio получает портфолио
func GetPortfolio(ch chan PortfolioData) {
	defer close(ch)

	files, err := file.GetFileAll("static/images/portfolio/")
	if err != nil {
		ch <- PortfolioData{nil, err}
	}

	ch <- PortfolioData{files, nil}
}
