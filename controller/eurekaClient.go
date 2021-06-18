package controller

import (
	"fmt"

	"github.com/ArthurHlt/go-eureka-client/eureka"
)

// Register Eureka Client
func ReigsterEurekaClient(hostname, appname, ip string, port int, ttl uint, isSsl bool) {
	client := eureka.NewClient([]string{
		"http://139.150.75.239:8761/eureka",
	})

	instance := eureka.NewInstanceInfo(hostname, appname, ip, port, ttl, isSsl)
	err := client.RegisterInstance("sms-service", instance)
	if err != nil {
		fmt.Println("[Eureka] client registry error : ", err)
	} else {
		fmt.Println("[Eureka] client registry success")
	}
}
