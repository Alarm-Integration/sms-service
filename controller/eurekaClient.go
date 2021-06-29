package controller

import (
	eurekaHandler "github.com/GreatLaboratory/go-sms/eureka-handler"
)

func ReigsterEurekaClient(defaultzone, app string, port int) error {

	client, clientCreateErr := eurekaHandler.NewClient(&eurekaHandler.Config{
		DefaultZone:           defaultzone,
		App:                   app,
		Port:                  port,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})

	if clientCreateErr != nil {
		return clientCreateErr
	}

	if clientStartErr := client.Start(); clientStartErr != nil {
		return clientStartErr
	}

	return nil

}
