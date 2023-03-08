package page_reg

import (
	"errors"
	"repair/pkg/database"
	"repair/pkg/hesh"
	ml "repair/pkg/mail"

	"github.com/doug-martin/goqu/v9"
)

// RegManager Создание база данных... Регистрация пользователя
func (m *Manager) RegManager() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	var login string
	ok, err := dialect.Select("login").From("auth").Where(goqu.C("login").In(m.Login)).ScanVal(&login)
	if err != nil {
		return err
	}

	if ok {
		return errors.New("менеджер с таким логином существует")
	}

	_, err = dialect.Insert("auth").Rows(goqu.Record{
		"login":      m.Login,
		"password":   hesh.NewHesh512(m.Password),
		"first_name": m.FirstName,
		"email":      m.Email,
		"admin":      m.Admin,
	}).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}

// MailSend отправка данных о регистрации
func (m *Manager) MailSend() error {
	if err := ml.Send(m.Email, "Доступ в систему", message(m.Login, m.Password, m.Email), nil); err != nil {
		return err
	}

	return nil
}

// RegStatusList Создание база данных... Регистрация список статусов
func (m *Manager) RegStatusList() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Insert("status").Cols("name", "translate").
		Vals(
			goqu.Vals{"Pause", "Пауза"},
			goqu.Vals{"Denied", "Отказано"},
			goqu.Vals{"To Do", "К выполнению"},
			goqu.Vals{"Done", "Готово"},
			goqu.Vals{"In Progress", "В процессе"},
		).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}

// RegIcon Регистрация списков иконок
func (m *Manager) RegIcon() error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dialect := database.Dialect(db)

	_, err = dialect.Insert("icon").Cols("name", "teg", "color").
		Vals(
			goqu.Vals{"Вконтакте", "fab fa-vk", "Vkontakte"},
			goqu.Vals{"Одноклассники", "fab fa-odnoklassniki", "Odnoklassniki"},
			goqu.Vals{"Facebook", "fab fa-facebook-f", "Facebook"},
			goqu.Vals{"Instagram", "fab fa-instagram", "Instagram"},
			goqu.Vals{"Telegram", "fab fa-telegram-plane", "Telegram"},
		).Executor().Exec()

	if err != nil {
		return err
	}
	return nil
}
