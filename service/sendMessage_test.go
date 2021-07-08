package service

import (
	"fmt"
	"github.com/GreatLaboratory/go-sms/model"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"time"
)

func Test_Send_Message_Success(t *testing.T) {
	defer gock.Off()

	Convey("Given right request body", t, func() {
		requestBody := model.RequestBody{
			Messages: []model.SendMessageDto{
				{
					To:   "01088350310",
					From: "01092988726",
					Text: "test",
					Type: "SMS",
				},
			},
		}
		var logValue []map[string]interface{}

		logValue = append(logValue, map[string]interface{}{
			"message": "메시지 그룹이 생성되었습니다.",
		})
		logValue = append(logValue, map[string]interface{}{
			"message": "단문문자(SMS) 1 건이 추가되었습니다.",
		})
		logValue = append(logValue, map[string]interface{}{
			"message": "메시지를 발송했습니다.",
		})

		responseBody := model.SendMessageResponseDto{
			Count: model.Count{Total: 1, SentTotal: 0, SentFailed: 0, SentSuccess: 0, SentPending: 0, SentReplacement: 0, Refund: 0, RegisteredFailed: 0, RegisteredSuccess: 1},
			Log:   logValue,
		}
		gock.New("http://api.coolsms.co.kr").
			Post("/messages/v4/send-many").
			JSON(requestBody).
			Reply(200).
			JSON(responseBody)

		Convey("When sending sms alarm", func() {
			err := SendMessage(requestBody)

			Convey("Then sms should be alarmed successfully", func() {
				So(err, ShouldBeNil)
				t.Log("메시지 그룹이 생성되었습니다.")
				t.Log("단문문자(SMS) 1 건이 추가되었습니다.")
				t.Log("메시지를 발송했습니다.")
				t.Log("1개의 알림발송 시도 중 성공 : 1 // 실패 : 0")
			})
		})
	})
}

func Test_Send_Message_Fail_By_Wrong_Number(t *testing.T) {
	defer gock.Off()

	Convey("Given string value at receiver number", t, func() {
		requestBody := model.RequestBody{
			Messages: []model.SendMessageDto{
				{
					To:   "test",
					From: "test",
					Text: "test",
					Type: "SMS",
				},
			},
		}
		var logValue []map[string]interface{}

		logValue = append(logValue, map[string]interface{}{
			"message": "메시지 그룹이 생성되었습니다.",
		})
		logValue = append(logValue, map[string]interface{}{
			"message": "단문문자(SMS) 1 건이 추가되었습니다.",
		})
		logValue = append(logValue, map[string]interface{}{
			"message": "메시지를 발송했습니다.",
		})

		responseBody := model.SendMessageResponseDto{
			Count: model.Count{Total: 1, SentTotal: 0, SentFailed: 0, SentSuccess: 0, SentPending: 0, SentReplacement: 0, Refund: 0, RegisteredFailed: 1, RegisteredSuccess: 0},
			Log:   logValue,
		}
		gock.New("http://api.coolsms.co.kr").
			Post("/messages/v4/send-many").
			JSON(requestBody).
			Reply(400).
			JSON(responseBody)

		Convey("When sending sms alarm", func() {
			err := SendMessage(requestBody)

			Convey("Then sms alarm should be failed", func() {
				So(err, ShouldBeNil)
				t.Log("메시지 그룹이 생성되었습니다.")
				t.Log("단문문자(SMS) 1 건이 추가되었습니다.")
				t.Log("메시지를 발송했습니다.")
				t.Log("1개의 알림발송 시도 중 성공 : 0 // 실패 : 1")
			})
		})
	})
}

func Test_Get_Authorization(t *testing.T) {

	Convey("Given parameters of function(getAuthorization)", t, func() {
		apiKey := "test"
		apiSecret := "test"
		date := time.Now().Format(time.RFC3339)
		salt := "b232b154adc18c017f73ee1e530ada8044230488"
		signature := "81878d9a148cff7c7cd873f61ac9e30a38e5f26ec95f612495caaed595c4ec0e"
		expectedAuthorization := fmt.Sprintf("HMAC-SHA256 apiKey=%s, date=%s, salt=%s, signature=%s", apiKey, date, salt, signature)

		Convey("When exec function(getAuthorization)", func() {
			actualAuthorization := getAuthorization(apiKey, apiSecret)

			Convey("Then authorization value should be not equal because of time difference", func() {
				So(actualAuthorization, ShouldNotEqual, expectedAuthorization)
			})
		})
	})
}

func Test_Random_String(t *testing.T) {

	Convey("Given size", t, func() {
		size := 50

		Convey("When making random string in three times", func() {
			salt1 := randomString(size)
			salt2 := randomString(size)
			salt3 := randomString(size)

			Convey("Then random string should not be equal with each other", func() {
				So(salt1, ShouldNotEqual, salt2)
				So(salt2, ShouldNotEqual, salt3)
				So(salt3, ShouldNotEqual, salt1)
			})
		})
	})
}
