package mail

type IMailer interface {
	Send(to string, subject string, body string, isHTML bool) error
}
