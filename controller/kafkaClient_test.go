package controller_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/GreatLaboratory/go-sms/controller"
	"github.com/stretchr/testify/assert"
)

func Test_Kafka_Connection_Success(t *testing.T) {

	// given
	kafkaServer := "139.150.75.240"
	groupId := "sms-service"
	topics := []string{"sms"}

	// when
	err := controller.ConnectKafkaConsumer(kafkaServer, groupId, topics, true)

	// then
	assert.Nil(t, err)
}

func Test_Kafka_Connection_Fail(t *testing.T) {

	// given
	kafkaServer := "1234.1234.1234"
	groupId := "sms-service"
	topics := []string{"sms"}
	expectedError := fmt.Sprintf("%s:9092/bootstrap: Failed to resolve '%s:9092':", kafkaServer, kafkaServer)

	// when
	err := controller.ConnectKafkaConsumer(kafkaServer, groupId, topics, true)

	// then
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)

}

func Test_Kafka_Consumer_Create_Fail(t *testing.T) {

	// given
	kafkaServer := "139.150.75.240"
	var groupId string
	topics := []string{"sms"}
	expectedErrorString := "consumer create failed"

	// when
	err := controller.ConnectKafkaConsumer(kafkaServer, groupId, topics, true)

	// then
	assert.NotNil(t, err)
	assert.EqualError(t, errors.New(expectedErrorString), expectedErrorString)

}

func Test_Kafka_Consumer_Subscribe_Fail(t *testing.T) {

	// given
	kafkaServer := "139.150.75.240"
	groupId := ""
	topics := []string{"sms"}
	expectedErrorString := "consumer topic subscribe failed"

	// when
	err := controller.ConnectKafkaConsumer(kafkaServer, groupId, topics, true)

	// then
	assert.NotNil(t, err)
	assert.EqualError(t, errors.New(expectedErrorString), expectedErrorString)

}
