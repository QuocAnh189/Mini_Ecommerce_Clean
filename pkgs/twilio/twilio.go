package twilio

import (
	"log"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type Client struct {
	accountSID      string
	authToken       string
	twilioNumber    string
	twilioServiceId string
}

type TwilioClient struct {
	info   Client
	client *twilio.RestClient
}

func NewTwilioSmsProvider(accountSid, authToken, twilioNumber, twilioServiceId string) *TwilioClient {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwilioClient{
		info: Client{
			accountSID:      accountSid,
			authToken:       authToken,
			twilioNumber:    twilioNumber,
			twilioServiceId: twilioServiceId,
		},
		client: client,
	}
}

func (c *TwilioClient) SendSMSByPhoneNumber(to, body string) error {
	log.Printf("SendSMSByPhoneNumber: Sending SMS to %s with from=%s\n", to, c.info.twilioNumber)

	params := buildParams(to, body)
	params.SetFrom(c.info.twilioNumber)
	_, err := c.client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	return nil
}

func (c *TwilioClient) SendSMSByMessagingService(to, body string) error {
	log.Printf("SendSMSByMessagingService: Sending SMS to %s with serviceId=%s\n", to, c.info.twilioServiceId)

	params := buildParams(to, body)
	params.SetMessagingServiceSid(c.info.twilioServiceId)
	_, err := c.client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	return nil
}

func buildParams(to, body string) *api.CreateMessageParams {
	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetBody(body)
	return params
}
