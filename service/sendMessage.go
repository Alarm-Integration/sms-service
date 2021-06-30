package service

import (
	"fmt"

	"github.com/GreatLaboratory/go-sms/model"
	coolsms "github.com/coolsms/coolsms-go"
	"github.com/spf13/viper"
)

// 메세지 발송
func SendGroupMessage(createGroupParams map[string]string, sendMessageDataList []model.SendMessageDto) error {
	client := createClient()

	// 1. 메세지 발송을 위한 그룹 생성
	groupId := createGroup(createGroupParams)

	// 2. 그룹에 메세지 데이터 저장
	addGroupMessage(groupId, sendMessageDataList)

	// 3. 완성된 그룹의 메세지 발송
	result, err := client.Messages.SendGroup(groupId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Print Result
	fmt.Println("SendGroupMessage result : ", result)
	return nil
}

// coolsms sdk에서 요청을 보내는 client 객체 생성
func createClient() *coolsms.Client {
	client := coolsms.NewClient()
	client.Messages.Config = map[string]string{
		"APIKey":    viper.GetString("sms.APIKey"),
		"APISecret": viper.GetString("sms.APISecret"),
		"Protocol":  viper.GetString("sms.Protocol"),
		"Domain":    viper.GetString("sms.Domain"),
	}
	return client
}

// 메세지 발송을 위한 그룹 생성
func createGroup(params map[string]string) string {
	client := createClient()

	createdGroup, err := client.Messages.CreateGroup(params)
	if err != nil {
		fmt.Println(err)
	}

	return createdGroup.GroupId
}

// 생성된 그룹에 메세지 데이터 저장
func addGroupMessage(groupId string, sendMessageDataList []model.SendMessageDto) {
	client := createClient()

	fmt.Println("message List : ", sendMessageDataList)

	params := make(map[string]interface{})
	params["messages"] = sendMessageDataList

	_, err := client.Messages.AddGroupMessage(groupId, params)
	if err != nil {
		fmt.Println(err)
		return
	}
}
