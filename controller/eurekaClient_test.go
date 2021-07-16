package controller

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Eureka_Registration_Success(t *testing.T) {

	Convey("Given", t, func() {
		defaultzone := "http://10.7.27.18:8761/eureka/"
		app := "sms-service"
		port := 30020

		Convey("When", func() {
			err := ReigsterEurekaClient(defaultzone, app, port)

			Convey("Then", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func Test_Eureka_Registration_Fail(t *testing.T) {

	Convey("Given", t, func() {
		defaultzone := "http://139.150.75.2391234:8761/eureka/"
		app := "sms-service"
		port := 30020
		expectedErr := "[Eureka] client registration failed"

		Convey("When", func() {
			err := ReigsterEurekaClient(defaultzone, app, port)

			Convey("Then", func() {
				So(err, ShouldBeError, expectedErr)
			})
		})
	})
}
