package mail

import (
	"os"
	"strconv"

	ml "github.com/go-mail/mail"
)

//Send отправляет письмо получателю
func Send(to, subject, body string, dirFileName interface{}) error {
	mess := ml.NewMessage()
	mess.SetHeader("From", os.Getenv("MAIL_LOGIN"))
	mess.SetHeader("To", to)
	mess.SetHeader("Subject", subject)
	mess.SetBody("text/html", body)

	if dirFileName != nil {
		mess.Attach(dirFileName.(string))
	}

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return err
	}

	d := ml.NewDialer(os.Getenv("MAIL_HOST"), port, os.Getenv("MAIL_LOGIN"), os.Getenv("MAIL_PASSWORD"))
	d.StartTLSPolicy = ml.MandatoryStartTLS

	if err = d.DialAndSend(mess); err != nil {
		return err
	}

	return nil
}
