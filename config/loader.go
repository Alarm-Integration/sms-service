package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GreatLaboratory/go-sms/model"
	"github.com/spf13/viper"
)

func LoadConfigurationFromBranch(configServerUrl string, appName string, profile string, branch string) {
	url := fmt.Sprintf("%s/%s/%s/%s", configServerUrl, appName, profile, branch)
	fmt.Printf("[Config] Loading config from %s\n", url)
	body, _ := fetchConfiguration(url)
	parseConfiguration(body)
}

func fetchConfiguration(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic("[Config] Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func parseConfiguration(body []byte) {
	var cloudConfig model.SpringCloudConfig
	err := json.Unmarshal(body, &cloudConfig)
	if err != nil {
		panic("[Config] Cannot parse configuration, message: " + err.Error())
	}

	for key, value := range cloudConfig.PropertySources[0].Source {
		viper.Set(key, value)
		fmt.Printf("[Config] Loading config property %v => %v\n", key, value)
	}
	if viper.IsSet("serviceName") {
		fmt.Printf("[Config] Successfully loaded configuration for service %s\n", viper.GetString("serviceName"))
	}
}
