package initial

import "ecommerce_clean/pkgs/mail"

func InitMailer(host string, port int, username, password, from string) *mail.Mailer {
	mailer := mail.NewMailer(host, port, username, password, from)
	return mailer
}
