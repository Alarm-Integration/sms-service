package main

import (
	"fmt"

	"github.com/GreatLaboratory/go-sms/controller"
)

func main() {

	// 1. Register Eureka Client to Discovery Service
	fmt.Println("[Eureka] Start Client Registration!!!")
	controller.ReigsterEurekaClient("http://139.150.75.239:8761/eureka/", "sms-service", 30020)

}
