package router

import (
	"repair/internal/page/admin/page_manager"
	"repair/internal/page/errs"
	"repair/internal/page/robots"

	"repair/pkg/session"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"repair/internal/page/install/page_form"
	"repair/internal/page/lending/page_back_mail"
	"repair/internal/page/lending/page_home"
	"repair/internal/page/lending/page_price_list"
	"repair/internal/page/lending/page_store"

	"repair/internal/page/admin/page_auth"
	"repair/internal/page/admin/page_cabinet"
	"repair/internal/page/admin/page_comment"
	"repair/internal/page/admin/page_filter"
	"repair/internal/page/admin/page_list_service"
	"repair/internal/page/admin/page_portfolio"
	"repair/internal/page/admin/page_setting"
	"repair/internal/page/admin/page_social_network"
	"repair/internal/page/admin/page_static"
	"repair/internal/page/admin/page_task"
	"repair/internal/page/admin/page_type_service"
)

// List роутер url адресов
func List(r *gin.Engine) {

	//___________________________Основной сайт________________________________

	// Основной сайт Лендинг для пользователя
	var lend = &page_home.Page{Router: r.Group("/")}
	lend.Home("")
	lend.Review("review")

	// Прайс-лист услуг
	var priceList = &page_price_list.Page{Router: lend.Router.Group("price-list")}
	priceList.PriceList("")
	priceList.PriceListSearch("/search")

	//Отправка заявки на обратную связь
	var backMail = &page_back_mail.Page{Router: lend.Router.Group("mail")}
	backMail.Send("")

	//Оформления услуги + обратную звязь с клиентом
	var store = &page_store.Page{Router: lend.Router.Group("store/:name")}
	store.Page("")
	store.Detail("/detail")
	store.Send("/detail/mail")

	//__________________________Панель админки (Личный кабинет) сайта_________________________

	//auth Авторизация
	var auth = &page_auth.Page{Router: lend.Router.Group("admin")}
	auth.Router.Use(sessions.Sessions("manager", session.StoreKeyGenerate()))

	auth.Page("")
	auth.Auth("/auth")
	auth.CheckToMail("/auth/check")
	auth.SessionExit("/exit")

	//cabinet Личный кабинет
	var cabinet = &page_cabinet.Page{Router: auth.Router.Group("cabinet")}
	cabinet.Page("")

	//setting Настройки
	var setting = &page_setting.Page{Router: cabinet.Router.Group("setting")}
	setting.Page("")
	setting.EditPassword("/password")
	setting.EditEmail("/email")
	setting.EditLogin("/login")
	setting.EditFullName("/full-name")

	//socialNetwork группы Соц Сетей
	var socialNetwork = &page_social_network.Page{Router: cabinet.Router.Group("social-network")}
	socialNetwork.Page("")
	socialNetwork.Detail("/detail/:id")
	socialNetwork.Add("/add")
	socialNetwork.Edit("/edit")
	socialNetwork.Delete("/delete")

	//portfolio Портфолио
	var portfolio = &page_portfolio.Page{Router: cabinet.Router.Group("portfolio")}
	portfolio.Page("")
	portfolio.Add("/add")
	portfolio.Delete("/delete")

	//typeService Категории ремонта
	var typeService = &page_type_service.Page{Router: cabinet.Router.Group("type-service")}
	typeService.Page("")
	typeService.Add("/add")
	typeService.Edit("/edit")
	typeService.EditPhoto("/editPhoto")
	typeService.Delete("/delete")
	typeService.Detail("/detail/:id")

	//listService Список услуг
	var listService = &page_list_service.Page{Router: cabinet.Router.Group("list-service")}
	listService.Page("")
	listService.Detail("/detail/:id")
	listService.Add("/add")
	listService.Edit("/edit")
	listService.EditPhoto("/editPhoto")
	listService.Delete("/delete")
	listService.Search("/search")

	//task задачи от заявок
	var task = &page_task.Page{Router: cabinet.Router.Group("task")}
	task.Page("")
	task.Detail("/:id")
	task.DetailEdit("/edit-detail/:id")
	task.EditStatus("/edit-status")
	task.EditPrice("/edit-price")
	task.EditManager("/edit-manager")
	task.EditTask("/edit")
	task.Delete("/delete")

	//comment добавляет коментарии к задачке
	var comment = &page_comment.Page{Router: cabinet.Router.Group("comment")}
	comment.Add("/add")
	comment.Detail("/detail/:id")
	comment.Edit("/edit")
	comment.EditPhoto("/edit-photo")
	comment.Delete("/delete")
	comment.DeletePhoto("/delete-phote")

	//static статистика
	var static = &page_static.Page{Router: cabinet.Router.Group("static")}
	static.Page("")

	//filter фильтр задач
	var filter = &page_filter.Page{Router: cabinet.Router.Group("filter")}
	filter.Page("")
	filter.Search("/task")

	//manager менеджеры
	var manager = &page_manager.Page{Router: cabinet.Router.Group("manager")}
	manager.Page("")
	manager.Detail("/detail/:id")
	manager.Add("/add")
	manager.EditPassword("/password")
	manager.EditEmail("/email")
	manager.EditLogin("/login")
	manager.EditFullName("/full-name")
	manager.EditType("/type")
	manager.Delete("/delete")

	//__________________________Настройка сайта______________________________

	//Установка + настройка панель админа
	var install = &page_form.Page{Router: lend.Router.Group("install")}
	install.Page("")
	install.Register("/reg")

	//Настрока для поисковых роботов
	var robot = &robots.Page{Router: lend.Router.Group("robots.txt")}
	robot.RobotsTXT("")

	//Обработчик ошибок
	r.NoRoute(errs.NotFound)
	r.NoMethod(errs.NotFound)
}
