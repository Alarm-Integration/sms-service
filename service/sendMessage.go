package service

import (
	"crypto/hmac"
	cr "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/GreatLaboratory/go-sms/model"
	"github.com/spf13/viper"
)

const uri string = "http://api.coolsms.co.kr/messages/v4/send-many"

func SendMessage(requestBody model.RequestBody) error {
	out, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	requestBodyString := string(out)
	data := strings.NewReader(requestBodyString)
	req, err := http.NewRequest("POST", uri, data)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", getAuthorization(viper.GetString("sms.APIKey"), viper.GetString("sms.APISecret")))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)

	return nil
}

func getAuthorization(apiKey string, apiSecret string) string {
	salt := randomString(20)
	date := time.Now().Format(time.RFC3339)
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(date + salt))
	signature := hex.EncodeToString(h.Sum(nil))
	authorization := fmt.Sprintf("HMAC-SHA256 apiKey=%s, date=%s, salt=%s, signature=%s", apiKey, date, salt, signature)
	return authorization
}

func randomString(n int) string {
	b := make([]byte, 20)
	_, _ = cr.Read(b)
	return hex.EncodeToString(b)
}

// // 메세지 발송
// func SendGroupMessage(createGroupParams map[string]string, sendMessageDataList []model.SendMessageDto) error {
// 	client := createClient()

// 	// 1. 메세지 발송을 위한 그룹 생성
// 	groupId, createGroupErr := createGroup(createGroupParams)
// 	if createGroupErr != nil {
// 		fmt.Println("[SMS] create group error")
// 		return createGroupErr
// 	}

// 	// 2. 그룹에 메세지 데이터 저장
// 	addMessageErr := addGroupMessage(groupId, sendMessageDataList)
// 	if addMessageErr != nil {
// 		fmt.Println("[SMS] add message error")
// 		return addMessageErr
// 	}

// 	// 3. 완성된 그룹의 메세지 발송
// 	_, sendMessageErr := client.Messages.SendGroup(groupId)
// 	if sendMessageErr != nil {
// 		fmt.Println("[SMS] send message error")
// 		return sendMessageErr
// 	}

// 	fmt.Println("[SMS] send message success")
// 	return nil
// }

// // 메세지를 보내는 client 객체 생성
// func createClient() *coolsms.Client {
// 	client := coolsms.NewClient()
// 	client.Messages.Config = map[string]string{
// 		"APIKey":    viper.GetString("sms.APIKey"),
// 		"APISecret": viper.GetString("sms.APISecret"),
// 		"Protocol":  viper.GetString("sms.Protocol"),
// 		"Domain":    viper.GetString("sms.Domain"),
// 		"Prefix":    "",
// 	}
// 	return client
// }

// // 메세지 발송을 위한 그룹 생성
// func createGroup(params map[string]string) (string, error) {
// 	client := createClient()
// 	createdGroup, err := client.Messages.CreateGroup(params)
// 	if err != nil {
// 		return "", err
// 	}

// 	fmt.Println("[SMS] create group success")
// 	return createdGroup.GroupId, nil
// }

// // 생성된 그룹에 메세지 데이터 저장
// func addGroupMessage(groupId string, sendMessageDataList []model.SendMessageDto) error {
// 	client := createClient()
// 	params := make(map[string]interface{})
// 	params["messages"] = sendMessageDataList

// 	_, err := client.Messages.AddGroupMessage(groupId, params)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("[SMS] add message success")
// 	return nil
// }
