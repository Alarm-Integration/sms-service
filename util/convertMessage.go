package util

import (
	"encoding/json"
	"fmt"

	"github.com/GreatLaboratory/go-sms/model"
)

func ConvertByteToDtoList(byteValue []byte) (model.RequestBody, string, error) {
	var requestBody model.RequestBody
	var requestId string
	var sendMessageDataList []model.SendMessageDto
	var topicMessageDto model.TopicMessageDto

	if err := json.Unmarshal(byteValue, &topicMessageDto); err != nil {
		return requestBody, requestId, err
	}

	text := fmt.Sprintf("%s\n%s", topicMessageDto.Title, topicMessageDto.Content)
	var messageType string
	if len(text) <= 90 {
		messageType = model.SMS.String()
	} else {
		messageType = model.LMS.String()
	}
	for _, address := range topicMessageDto.Addresses {
		sendMessageData := model.SendMessageDto{
			To:   address,
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
	for i, v := range topicMessageDto.Addresses {
		fmt.Println("address", i, " : ", v)
	}
	fmt.Println("=====================================================")

	requestBody.Messages = sendMessageDataList
	requestId = topicMessageDto.TraceID
	return requestBody, requestId, nil
}
