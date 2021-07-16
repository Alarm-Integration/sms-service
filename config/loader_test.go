package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func Test_Parse_Configuration_Success(t *testing.T) {

	Convey("Given a JSON configuration response body", t, func() {
		var body = `{
			"name":"sms-service",
			"profiles":["test"],
			"label":null,
			"version":null,
			"propertySources":[
				{
					"name":"file:/config-repo/sms-service.yml",
					"source":{
						"serviceName":"sms-service",
						"server_port":6767,
						"server_name":"sms-service test"
					}
				}
			]
		}`

		Convey("When parsed", func() {
			parseConfiguration([]byte(body))

			Convey("Then Viper should have been populated with values from Source", func() {
				So(viper.GetString("server_name"), ShouldEqual, "sms-service test")
				So(viper.GetString("server_port"), ShouldEqual, "6767")
				So(viper.IsSet("serviceName"), ShouldBeTrue)
			})
		})
	})
}

func Test_Parse_Configuration_Fail(t *testing.T) {

	Convey("Given a wrong JSON configuration response body", t, func() {
		var body = `{wrongwrongwrong
			"name":"sms-service",
			"profiles":["test"],
			"label":null,
			"version":null,
			"propertySources":[
				{
					"name":"file:/config-repo/sms-service.yml",
					"source":{
						"server_port":6767,
						"server_name":"sms-service test",
					}
				}
			]
		}`

		Convey("When parsed", func() {
			panicFunc := func() { parseConfiguration([]byte(body)) }

			Convey("Then parseConfiguration function should panic with unmarshal error", func() {
				So(panicFunc, ShouldPanicWith, "[Config] Cannot parse configuration, message: invalid character 'w' looking for beginning of object key string")
			})
		})
	})
}

func Test_Fetch_Configuration_Fail(t *testing.T) {

	Convey("Given a wrong url", t, func() {
		url := "wrongUrl"

		Convey("When parsed", func() {
			panicFunc := func() { fetchConfiguration(url) }

			Convey("Then fetchConfiguration function should panic with wrong url error", func() {
				So(panicFunc, ShouldPanicWith, `[Config] Couldn't load configuration, cannot start. Terminating. Error: Get "wrongUrl": unsupported protocol scheme ""`)
			})
		})
	})
}
