package initial

import (
	"ecommerce_clean/pkgs/twilio"
)

func InitTwilioProvider(accountSid, authToken, twilioNumber, twilioServiceId string) (*twilio.TwilioClient, error) {
	// twilioProvider := twilio.NewTwilioSmsProvider(accountSid, authToken, twilioNumber, twilioServiceId)
	// if twilioProvider == nil {
	// 	return nil, fmt.Errorf("failed to initialize Twilio provider")
	// }

	// logger.Info("Twilio provider initialized successfully")
	// return twilioProvider, nil
	fake := &twilio.TwilioClient{}
	return fake, nil
}
