package service

import (
	"crypto/hmac"
	cr "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/GreatLaboratory/go-sms/util"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/GreatLaboratory/go-sms/model"
	"github.com/spf13/viper"
)

const URI string = "http://api.coolsms.co.kr/messages/v4/send-many"

func SendMessage(requestBody model.RequestBody, alarmResultLog model.AlarmResultLogDto) error {
	out, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	requestBodyString := string(out)
	data := strings.NewReader(requestBodyString)
	req, err := http.NewRequest("POST", URI, data)
	if err != nil {
		return err
	}

	authorization := getAuthorization(viper.GetString("sms.APIKey"), viper.GetString("sms.APISecret"))
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	var response model.SendMessageResponseDto
	if err := json.Unmarshal(bytes, &response); err != nil {
		return err
	}

	for _, logValue := range response.Log {
		fmt.Println(logValue["message"])
	}

	fmt.Printf("%v개의 알림발송 시도 중 성공 : %v // 실패 : %v", response.Count.Total, response.Count.RegisteredSuccess, response.Count.RegisteredFailed)

	resultMsg := fmt.Sprintf("%v개의 알림발송 시도 중 성공 : %v // 실패 : %v", response.Count.Total, response.Count.RegisteredSuccess, response.Count.RegisteredFailed)
	userID := alarmResultLog.UserID
	traceID := alarmResultLog.TraceID

	err = util.FluentdSender(userID, traceID, resultMsg)
	if err != nil {
		return err
	}

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
	b := make([]byte, n)
	_, _ = cr.Read(b)
	return hex.EncodeToString(b)
}
