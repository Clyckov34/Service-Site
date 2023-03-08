package page_back_mail

import (
	"log"
	ml "repair/pkg/mail"
	"repair/pkg/manager"
)

type Param struct {
	FirstName string
	Phone     string
	Email     string
	Text      string
}

// Send отправка на почту
func (m *Param) Send() error {
	listEmail, err := manager.GetEmailTypeAdmin()
	if err != nil {
		return err
	}

	sendMessage := func(email string) {
		if err := ml.Send(email, "Обратная связь", Message(m), nil); err != nil {
			log.Println(err)
		}
	}

	for _, v := range listEmail {
		go sendMessage(v)
	}

	return nil
}
