package main

import (
	"github.com/GreatLaboratory/go-sms/controller"
)

func main() {
	// 1. Eureka Client Register
	controller.ReigsterEurekaClient("alarm-integration.com:sms-service:8081", "sms-service", "discovery-service", 8081, 80, false)
}
