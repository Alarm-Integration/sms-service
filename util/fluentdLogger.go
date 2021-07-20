package util

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/spf13/viper"
)

const Tag string = "alarm.result.access"
const FluentdPort int = 24224

func FluentdSender(isSuccess bool, address, requestID, logMessage string) error {
	logger, err := createLogger()
	if err != nil {
		return err
	}
	defer logger.Close()

	var data = map[string]interface{}{
		"request_id":  requestID,
		"app_name":    "sms",
		"log_message": logMessage,
		"is_success":  isSuccess,
		"address":     address,
	}
	err = send(logger, data)
	if err != nil {
		return err
	}

	return nil
}

func createLogger() (*fluent.Fluent, error) {
	logger, err := fluent.New(fluent.Config{
		FluentHost: viper.GetString("fluentd.server"),
		FluentPort: FluentdPort,
	})
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func send(logger *fluent.Fluent, data map[string]interface{}) error {
	err := logger.Post(Tag, data)
	if err != nil {
		return err
	}

	return nil
}
