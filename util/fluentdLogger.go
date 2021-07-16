package util

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/spf13/viper"
)

const Tag string = "alarm.access"
const FluentdPort int = 24224

func FluentdSender(userID int, traceID, resultMsg string) error {
	logger, err := createLogger()
	if err != nil {
		return err
	}
	defer logger.Close()

	var data = map[string]interface{}{
		"user_id":    userID,
		"trace_id":   traceID,
		"app_name":   "sms",
		"result_msg": resultMsg,
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
