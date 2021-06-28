package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GreatLaboratory/go-sms/controller"
)

func main() {

	// 0. If env variables are not defined, terminate app
	if os.Getenv("EUREKA_SERVER") == "" {
		os.Exit(-1)
	}
	if os.Getenv("SMS_SERVICE_PORT") == "" {
		os.Exit(-1)
	}
	if os.Getenv("SERVICE_NAME") == "" {
		os.Exit(-1)
	}

	// 1. Register Eureka Client to Discovery Service
	fmt.Println("[Eureka] Start Client Registration!!!")
	port, _ := strconv.Atoi(os.Getenv("SMS_SERVICE_PORT"))
	controller.ReigsterEurekaClient(os.Getenv("EUREKA_SERVER"), os.Getenv("SERVICE_NAME"), port)

}
