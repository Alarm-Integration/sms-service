package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GreatLaboratory/go-sms/controller"
)

func main() {

	// 1. Register Eureka Client to Discovery Service
	fmt.Println("[Eureka] Start Client Registration!!!")
	port, err := strconv.Atoi(os.Getenv("SMS_SERVICE_PORT"))
	if err != nil {
		port = 30020
	}
	controller.ReigsterEurekaClient("http://139.150.75.239:8761/eureka/", "sms-service", port)

}
