package page_social_network

import (
	"net/http"
	"repair/pkg/check"
	"repair/pkg/errs"
	"repair/pkg/session"

	"repair/service/admin/page_social_network"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Router *gin.RouterGroup
}

// Page социальные сети форма
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

		var socialNetwork = &page_social_network.Param{}

		dataNetwork := make(chan page_social_network.ListNetworkAllChan)
		dataIcon := make(chan page_social_network.IconChan)

		go socialNetwork.GetAll(dataNetwork)
		go socialNetwork.GetIcon(dataIcon)

		network := <-dataNetwork
		icon := <-dataIcon

		if network.Err != nil {
			dataNetwork = nil
		}

		if icon.Err != nil {
			errs.HTTPStatusJson(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_social_network.html", gin.H{
			"Group": network.Data,
			"Icon":  icon.Data,
		})

	})

}

// Detail выводит форму для редактирование
func (m *Page) Detail(url string) {
	m.Router.GET(url, func(ctx *gin.Context) {

		sess, err := session.Check(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin")
			return
		}

		param, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, "Неверный формат ID")
			return
		}

		if err := check.ManagerAdmin(sess.Admin.(bool)); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		var socialNetwork = &page_social_network.Param{Id: param}

		dataEdit := make(chan page_social_network.ListNetworkIdChan)
		network := make(chan page_social_network.IconChan)

		go socialNetwork.GetID(dataEdit)
		go socialNetwork.GetIcon(network)

		data := <-dataEdit
		icon := <-network

		if data.Err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		if icon.Err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "cabinet_social_network_detail.html", gin.H{
			"Icon": icon.Data,
			"Id":   data.Data.Id,
			"Url":  data.Data.Url,
		})
	})
}

// Add добавляет групу соц сети
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

		icon, err := strconv.Atoi(ctx.PostForm("Icon"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		var socialNetwork = &page_social_network.Param{
			Url:    strings.TrimSpace(ctx.PostForm("Url")),
			IdIcon: icon,
		}

		if err := socialNetwork.Add(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/social-network")
	})

}

// Edit редактирует записи о соц сети
func (m *Page) Edit(url string) {
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

		icon, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Icon")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, "Неверный формат ID")
			return
		}

		id, err := strconv.Atoi(strings.TrimSpace(ctx.PostForm("Id")))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, "Неверный формат ID")
			return
		}

		var socialNetwork = &page_social_network.Param{
			Id:     id,
			IdIcon: icon,
			Url:    strings.TrimSpace(ctx.PostForm("Url")),
		}

		if err := socialNetwork.Edit(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/social-network")

	})

}

// Delete удаляет соц сеть
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

		id, err := strconv.Atoi(ctx.PostForm("Id"))
		if err != nil {
			errs.HTTPStatusJson(ctx, http.StatusBadRequest, "Неверный формат ID")
			return
		}

		var socialNetwork = &page_social_network.Param{Id: id}

		if err := socialNetwork.Delete(); err != nil {
			errs.HTTPStatusJson(ctx, http.StatusForbidden, err.Error())
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, "/admin/cabinet/social-network")

	})

}
