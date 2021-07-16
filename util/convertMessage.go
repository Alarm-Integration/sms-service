package util

import (
	"encoding/json"
	"fmt"

	"github.com/GreatLaboratory/go-sms/model"
)

func ConvertByteToDtoList(byteValue []byte) (model.RequestBody, model.AlarmResultLogDto, error) {
	var requestBody model.RequestBody
	var alarmResultLog model.AlarmResultLogDto
	var sendMessageDataList []model.SendMessageDto
	var topicMessageDto model.TopicMessageDto

	if err := json.Unmarshal(byteValue, &topicMessageDto); err != nil {
		return requestBody, alarmResultLog, err
	}

	text := fmt.Sprintf("%s\n%s", topicMessageDto.Title, topicMessageDto.Content)
	var messageType string
	if len(text) <= 90 {
		messageType = model.SMS.String()
	} else {
		messageType = model.LMS.String()
	}
	for _, receiver := range topicMessageDto.Receivers {
		sendMessageData := model.SendMessageDto{
			To:   receiver,
			From: "01092988726",
			Text: text,
			Type: messageType,
		}
		sendMessageDataList = append(sendMessageDataList, sendMessageData)
	}

	fmt.Println("=====================================================")
	fmt.Println("title : ", topicMessageDto.Title)
	fmt.Println("content : ", topicMessageDto.Content)
	fmt.Println("traceId : ", topicMessageDto.TraceID)
	fmt.Println("userId : ", topicMessageDto.UserID)
	for i, v := range topicMessageDto.Receivers {
		fmt.Println("receiver", i, " : ", v)
	}
	fmt.Println("=====================================================")

	requestBody.Messages = sendMessageDataList
	alarmResultLog = model.AlarmResultLogDto{
		UserID:  topicMessageDto.UserID,
		TraceID: topicMessageDto.TraceID,
	}
	return requestBody, alarmResultLog, nil
}
