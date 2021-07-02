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
	groupId, createGroupErr := createGroup(createGroupParams)
	if createGroupErr != nil {
		fmt.Println("[SMS] create group error")
		return createGroupErr
	}

	// 2. 그룹에 메세지 데이터 저장
	addMessageErr := addGroupMessage(groupId, sendMessageDataList)
	if addMessageErr != nil {
		fmt.Println("[SMS] add message error")
		return addMessageErr
	}

	// 3. 완성된 그룹의 메세지 발송
	_, sendMessageErr := client.Messages.SendGroup(groupId)
	if sendMessageErr != nil {
		fmt.Println("[SMS] send message error")
		return sendMessageErr
	}

	fmt.Println("[SMS] send message success")
	return nil
}

// 메세지를 보내는 client 객체 생성
func createClient() *coolsms.Client {
	client := coolsms.NewClient()
	client.Messages.Config = map[string]string{
		"APIKey":    viper.GetString("sms.APIKey"),
		"APISecret": viper.GetString("sms.APISecret"),
		"Protocol":  viper.GetString("sms.Protocol"),
		"Domain":    viper.GetString("sms.Domain"),
		"Prefix":    "",
	}
	return client
}

// 메세지 발송을 위한 그룹 생성
func createGroup(params map[string]string) (string, error) {
	client := createClient()
	createdGroup, err := client.Messages.CreateGroup(params)
	if err != nil {
		return "", err
	}

	fmt.Println("[SMS] create group success")
	return createdGroup.GroupId, nil
}

// 생성된 그룹에 메세지 데이터 저장
func addGroupMessage(groupId string, sendMessageDataList []model.SendMessageDto) error {
	client := createClient()
	params := make(map[string]interface{})
	params["messages"] = sendMessageDataList

	_, err := client.Messages.AddGroupMessage(groupId, params)
	if err != nil {
		return err
	}

	fmt.Println("[SMS] add message success")
	return nil
}
