package util

import (
	"github.com/GreatLaboratory/go-sms/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Convert_Message_Success_SMS(t *testing.T) {

	Convey("Given topic message byte received (SMS)", t, func() {
		topicMessageByte := []byte(`{
										"userId":30000000000,
										"traceId":"1",
										"groupId":null,
										"raws":["01092988726","01023255906"],
										"title":"Hello?",
										"content":"안녕하세요 좋은 발표 시간입니다."
                        			}`)
		expectedDtoList := model.RequestBody{
			Messages: []model.SendMessageDto{
				{
					To:   "01092988726",
					From: "01092988726",
					Text: "Hello?\n안녕하세요 좋은 발표 시간입니다.",
					Type: "SMS",
				}, {
					To:   "01023255906",
					From: "01092988726",
					Text: "Hello?\n안녕하세요 좋은 발표 시간입니다.",
					Type: "SMS",
				},
			},
		}

		Convey("When convert byte array to sendMessageDtoList", func() {
			actualDtoList, err := ConvertByteToDtoList(topicMessageByte)

			Convey("Then topic message byte converted", func() {
				So(err, ShouldBeNil)
				So(actualDtoList, ShouldResemble, expectedDtoList)
				So(actualDtoList, ShouldHaveSameTypeAs, expectedDtoList)
			})
		})
	})
}

func Test_Convert_Message_Success_LMS(t *testing.T) {

	Convey("Given topic message byte received (LMS)", t, func() {
		topicMessageByte := []byte(`{
										"userId":30000000000,
										"traceId":"1",
										"groupId":null,
										"raws":["01092988726"],
										"title":"Hello?",
										"content":"안녕하세요 좋은 발표 시간입니다.안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다.안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다.안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다."
                        			}`)
		expectedDtoList := model.RequestBody{
			Messages: []model.SendMessageDto{
				{
					To:   "01092988726",
					From: "01092988726",
					Text: "Hello?\n안녕하세요 좋은 발표 시간입니다.안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다.안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다.안녕하세요 좋은 발표 시간입니다. 안녕하세요 좋은 발표 시간입니다.",
					Type: "LMS",
				},
			},
		}

		Convey("When convert byte array to sendMessageDtoList", func() {
			actualDtoList, err := ConvertByteToDtoList(topicMessageByte)

			Convey("Then topic message byte converted", func() {
				So(err, ShouldBeNil)
				So(actualDtoList, ShouldResemble, expectedDtoList)
				So(actualDtoList, ShouldHaveSameTypeAs, expectedDtoList)
			})
		})
	})
}

func Test_Convert_Message_Fail(t *testing.T) {

	Convey("Given wrong topic message byte received", t, func() {
		topicMessageByte := []byte("byte-test")
		expectedDtoList := model.RequestBody{Messages: []model.SendMessageDto(nil)}

		Convey("When convert byte array to sendMessageDtoList", func() {
			actualDtoList, err := ConvertByteToDtoList(topicMessageByte)

			Convey("Then a refresh event received", func() {
				So(actualDtoList, ShouldResemble, expectedDtoList)
				So(err, ShouldBeError, "invalid character 'b' looking for beginning of value")
			})
		})
	})
}
