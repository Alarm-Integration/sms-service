package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/GreatLaboratory/go-sms/config"
	"github.com/GreatLaboratory/go-sms/controller"
	"github.com/spf13/viper"
)

// init function, runs before main()
func init() {

	// If env variables are not defined, terminate app
	if os.Getenv("SERVICE_NAME") == "" ||
		os.Getenv("SMS_SERVICE_PORT") == "" ||
		os.Getenv("CONFIG_SERVER") == "" {
		os.Exit(-1)
	}

	configServerUrl := flag.String("configServerUrl", os.Getenv("CONFIG_SERVER"), "Address to config server")
	serviceName := flag.String("serviceName", os.Getenv("SERVICE_NAME"), "service name of this application")
	servicePort := flag.String("servicePort", os.Getenv("SMS_SERVICE_PORT"), "service port of this application")
	profile := flag.String("profile", "default", "Environment profile, something similar to spring profiles")
	configBranch := flag.String("configBranch", "master", "git branch to fetch configuration from")

	flag.Parse()

	fmt.Println("Specified configBranch is " + *configBranch)
	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
	viper.Set("serviceName", *serviceName)
	viper.Set("servicePort", *servicePort)
}

func main() {

	// 0. load the config
	config.LoadConfigurationFromBranch(
		viper.GetString("configserverurl"),
		viper.GetString("serviceName"),
		viper.GetString("profile"),
		viper.GetString("configbranch"),
	)
	amqpServer := fmt.Sprintf("amqp://%s:%s@%s:%s", viper.GetString("rabbitmq.username"), viper.GetString("rabbitmq.password"), viper.GetString("rabbitmq.server"), viper.GetString("rabbitmq.port"))
	go config.StartListener(amqpServer, "springCloudBus", "topic", "sms-service-queue", "springCloudBus", viper.GetString("serviceName"))

	// // 1. Register Eureka Client to Discovery Service
	fmt.Println("[Eureka] Start Client Registration!!!")
	port, _ := strconv.Atoi(viper.GetString("server.port"))
	controller.ReigsterEurekaClient(viper.GetString("eureka.server"), viper.GetString("serviceName"), port)

	// // 2. Connect to Kafka Broker (if consuming message, send sms alarm)
	fmt.Println("[KAFKA] Start Connection!!!")
	controller.ConnectKafkaConsumer(viper.GetString("kafka.server"), viper.GetString("serviceName"), []string{"sms"})
}
