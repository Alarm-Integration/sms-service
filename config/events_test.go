package config

import (
	"os"
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gock.v1"
)

const SERVICE_NAME = "sms-service"

func Test_Handle_Refresh_Event_Success(t *testing.T) {
	// Configure initial viper values
	viper.Set("configServerUrl", "http://config-server:8888")
	viper.Set("profile", "test")
	viper.Set("configBranch", "master")

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
		body := `{
			"type":"RefreshRemoteApplicationEvent",
			"timestamp":1494514362123,
			"originService":"config-server:docker:8888",
			"destinationService":"sms-service:**",
			"id":"53e61c71-cbae-4b6d-84bb-d0dcc0aeb4dc"
		}`
		Convey("When handled", func() {
			handleRefreshEvent([]byte(body), SERVICE_NAME)

			Convey("Then Viper should have been re-populated with values from Source", func() {
				So(viper.GetString("server_port"), ShouldEqual, "6767")
				So(viper.GetString("server_name"), ShouldEqual, "SMS-SERVICE RELOADED")
			})
		})
	})
}

func Test_Handle_Refresh_Event_Fail(t *testing.T) {
	// Configure initial viper values
	viper.Set("configServerUrl", "http://config-server:8888")
	viper.Set("profile", "test")
	viper.Set("configBranch", "master")

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

	Convey("Given a wrong refresh event received", t, func() {
		body := `{
			wrongwrongwrongwrong
			"type":"RefreshRemoteApplicationEvent",
			"timestamp":1494514362123,
			"originService":"config-server:docker:8888",
			"destinationService":"sms-service:**",
			"id":"53e61c71-cbae-4b6d-84bb-d0dcc0aeb4dc"
		}`
		Convey("When handled", func() {
			panicFunc := func() { handleRefreshEvent([]byte(body), SERVICE_NAME) }

			Convey("Then handleRefreshEvent function should panic with unmarshal error", func() {
				So(panicFunc, ShouldPanicWith, "[Config] Problem parsing UpdateToken: %vinvalid character 'w' looking for beginning of object key string")
			})
		})
	})
}

func Test_StartListener_Fail(t *testing.T) {

	Convey("Given", t, func() {
		amqpURI := "amqp://test-server:5762"
		exchange := "test"
		exchangeType := "topic"
		queueName := "test-queue"
		bindingKey := "springCloudBusTest"
		consumerTag := "test"
		expectedErr := "exit status 1"

		Convey("When", func() {
			if os.Getenv("FLAG") == "1" {
				StartListener(amqpURI, exchange, exchangeType, queueName, bindingKey, consumerTag)
				return
			}
			// Run the test in a subprocess
			cmd := exec.Command(os.Args[0], "-test.run=Test_StartListener_Fail")
			cmd.Env = append(os.Environ(), "FLAG=1")
			err := cmd.Run()
			err, ok := err.(*exec.ExitError)

			Convey("Then", func() {
				// Cast the error as *exec.ExitError and compare the result
				So(ok, ShouldBeTrue)
				So(err, ShouldBeError, expectedErr)
			})
		})
	})
}
