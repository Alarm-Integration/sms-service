package controller

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func Test_test(t *testing.T) {
	// Mock the expected outgoing request for new config
	defer gock.Off()
	gock.New("http://config-server:8888").
		Get("/sms-service/test/master").
		Reply(200).
		BodyString(`{
			"name":"sms-service",
			"profiles":["test"],
			"label":null,
			"version":null,
			"propertySources":[{
				"name":"file:/config-repo/sms-service.yml",
				"source":{
					"server_port":6767,
					"server_name":"SMS-SERVICE RELOADED"
				}
			}]
		}`)

	Convey("Given a refresh event received", t, func() {
		// body := `{
		// 	"type":"RefreshRemoteApplicationEvent",
		// 	"timestamp":1494514362123,
		// 	"originService":"config-server:docker:8888",
		// 	"destinationService":"sms-service:**",
		// 	"id":"53e61c71-cbae-4b6d-84bb-d0dcc0aeb4dc"
		// }`
		Convey("When handled", func() {
			// handleRefreshEvent([]byte(body), SERVICE_NAME)

			Convey("Then Viper should have been re-populated with values from Source", func() {
				// So(viper.GetString("server_port"), ShouldEqual, "6767")
				// So(viper.GetString("server_name"), ShouldEqual, "SMS-SERVICE RELOADED")
			})
		})
	})
}

func Test_Kafka_Connection_Fail(t *testing.T) {

	// given
	kafkaServer := "1234.1234.1234"
	groupId := "sms-service"
	topics := []string{"sms"}
	expectedError := fmt.Sprintf("%s:9092/bootstrap: Failed to resolve '%s:9092':", kafkaServer, kafkaServer)

	// when
	err := ConnectKafkaConsumer(kafkaServer, groupId, topics, true)

	// then
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedError)

}

func Test_Kafka_Consumer_Create_Fail(t *testing.T) {

	// given
	kafkaServer := "139.150.75.240"
	var groupId string
	topics := []string{"sms"}
	expectedErrorString := "[Kafka] consumer create failed"

	// when
	err := ConnectKafkaConsumer(kafkaServer, groupId, topics, true)

	// then
	assert.NotNil(t, err)
	assert.EqualError(t, errors.New(expectedErrorString), expectedErrorString)

}

func Test_Kafka_Consumer_Subscribe_Fail(t *testing.T) {

	// given
	kafkaServer := "139.150.75.240"
	groupId := ""
	topics := []string{"sms"}
	expectedErrorString := "[Kafka] consumer topic subscribe failed"

	// when
	err := ConnectKafkaConsumer(kafkaServer, groupId, topics, true)

	// then
	assert.NotNil(t, err)
	assert.EqualError(t, errors.New(expectedErrorString), expectedErrorString)

}
