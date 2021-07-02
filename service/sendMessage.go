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
