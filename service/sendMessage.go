package service

import (
	"crypto/hmac"
	cr "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/GreatLaboratory/go-sms/model"
	"github.com/spf13/viper"
)

const URI string = "http://api.coolsms.co.kr/messages/v4/send"

var CH = make(chan *model.ResultLogDto)

func SendMessage(sendMessageDto model.SendMessageDto, requestID string) {
	requestBody := model.SendRequestDto{
		Message: sendMessageDto,
	}
	out, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf(err.Error())
	}

	requestBodyString := string(out)
	data := strings.NewReader(requestBodyString)
	req, err := http.NewRequest("POST", URI, data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	authorization := getAuthorization(viper.GetString("sms.APIKey"), viper.GetString("sms.APISecret"))
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	var successResponse model.SendMessageSuccessResponseDto
	var failResponse model.SendMessageFailResponseDto

	var isSuccess bool
	var logMessage string

	if resp.StatusCode != 200 {
		if err := json.Unmarshal(bytes, &failResponse); err != nil {
			log.Fatalf(err.Error())
		}
		isSuccess = false
		logMessage = failResponse.ErrorMessage
		log.Printf("발송 실패 ::: %s", logMessage)
	} else {
		if err := json.Unmarshal(bytes, &successResponse); err != nil {
			log.Fatalf(err.Error())
		}
		isSuccess = true
		logMessage = successResponse.StatusMessage
		log.Printf("발송 성공 ::: %s", logMessage)
	}

	CH <- &model.ResultLogDto{
		IsSuccess:  isSuccess,
		Address:    sendMessageDto.To,
		RequestID:  requestID,
		LogMessage: logMessage,
	}

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
	b := make([]byte, n)
	_, _ = cr.Read(b)
	return hex.EncodeToString(b)
}
