package controller_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/GreatLaboratory/go-sms/controller"
	"github.com/stretchr/testify/assert"
)

func Test_Eureka_Registration_Success(t *testing.T) {

	// given
	defaultzone := os.Getenv("EUREKA_SERVER")
	app := "sms-service"
	port, portErr := strconv.Atoi(os.Getenv("SMS_SERVICE_PORT"))
	if portErr != nil {
		port = 30020
	}

	// when
	err := controller.ReigsterEurekaClient(defaultzone, app, port)

	// then
	assert.Nil(t, err)

}

func Test_Eureka_Registration_Fail(t *testing.T) {

	// given
	defaultzone := "http://139.150.75.2391234:8761/eureka/"
	app := "sms-service"
	port, portErr := strconv.Atoi(os.Getenv("SMS_SERVICE_PORT"))
	if portErr != nil {
		port = 30020
	}
	expectedErrorString := "client registration failed"

	// when
	err := controller.ReigsterEurekaClient(defaultzone, app, port)

	// then
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorString)

}
