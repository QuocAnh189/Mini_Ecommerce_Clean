package twilio

type Provider interface {
	SendSMSByPhoneNumber(to, body string) error
	SendSMSByMessagingService(to, body string) error
}
