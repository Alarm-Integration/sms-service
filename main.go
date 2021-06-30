package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GreatLaboratory/go-sms/controller"
)

func main() {

	// 0. If env variables are not defined, terminate app
	if os.Getenv("EUREKA_SERVER") == "" ||
		os.Getenv("KAFKA_SERVER") == "" ||
		os.Getenv("SMS_SERVICE_PORT") == "" ||
		os.Getenv("SERVICE_NAME") == "" {
		os.Exit(-1)
	}

	// 1. Register Eureka Client to Discovery Service
	fmt.Println("[Eureka] Start Client Registration!!!")
	port, _ := strconv.Atoi(os.Getenv("SMS_SERVICE_PORT"))
	controller.ReigsterEurekaClient(os.Getenv("EUREKA_SERVER"), os.Getenv("SERVICE_NAME"), port)

	// 2. Connect to Kafka Broker (if consuming message, send sms alarm)
	fmt.Println("[KAFKA] Start Connection!!!")
	controller.ConnectKafkaConsumer(os.Getenv("EUREKA_SERVER"), os.Getenv("SERVICE_NAME"), []string{"sms"})
}
