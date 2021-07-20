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
		requestBody := model.SendMessageDto{
			To:   "01088350310",
			From: "01092988726",
			Text: "test",
			Type: "SMS",
		}
		requestID := "test_request_id"

		responseBody := model.SendMessageSuccessResponseDto{
			GroupId:       "test_group_id",
			MessageId:     "test_message_id",
			AccountId:     "test_account_id",
			StatusCode:    "200",
			StatusMessage: "success",
			Country:       "82+",
			Type:          "SMS",
			To:            "01012341234",
			From:          "01012341234",
		}

		gock.New("http://api.coolsms.co.kr").
			Post("/messages/v4/send").
			JSON(requestBody).
			Reply(200).
			JSON(responseBody)

		Convey("When sending sms alarm", func() {
			SendMessage(requestBody, requestID)

			Convey("Then sms should be alarmed successfully", func() {
				t.Logf("발송 성공 ::: %s", responseBody.StatusMessage)
			})
		})
	})
}

func Test_Send_Message_Fail_By_Wrong_Number(t *testing.T) {
	defer gock.Off()

	Convey("Given string value at receiver number", t, func() {
		requestBody := model.SendMessageDto{
			To:   "01088350310",
			From: "01092988726",
			Text: "test",
			Type: "SMS",
		}
		requestID := "test_request_id"

		responseBody := model.SendMessageFailResponseDto{
			ErrorCode:    "400",
			ErrorMessage: "ValidationError",
		}

		gock.New("http://api.coolsms.co.kr").
			Post("/messages/v4/send").
			JSON(requestBody).
			Reply(400).
			JSON(responseBody)

		Convey("When sending sms alarm", func() {
			SendMessage(requestBody, requestID)

			Convey("Then sms alarm should be failed", func() {
				t.Logf("발송 실패 ::: %s", responseBody.ErrorMessage)
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
