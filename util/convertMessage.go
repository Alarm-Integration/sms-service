package util

import (
	"encoding/json"
	"fmt"

	"github.com/GreatLaboratory/go-sms/model"
)

func ConvertByteToDtoList(byteValue []byte) ([]model.SendMessageDto, error) {

	var sendMessageDataList []model.SendMessageDto
	var topicMessageDto model.TopicMessageDto

	if err := json.Unmarshal(byteValue, &topicMessageDto); err != nil {
		return nil, err
	}

	text := "[" + topicMessageDto.Title + "] \n" + topicMessageDto.Content
	var messageType string
	if len(text) <= 90 {
		messageType = model.SMS.String()
	} else {
		messageType = model.LMS.String()
	}
	for _, receiver := range topicMessageDto.Raws {
		sendMessageData := model.SendMessageDto{
			To:   receiver,
			From: topicMessageDto.Sender,
			Text: text,
			Type: messageType,
		}
		sendMessageDataList = append(sendMessageDataList, sendMessageData)
	}

	fmt.Println("=====================================================")
	fmt.Println("title : ", topicMessageDto.Title)
	fmt.Println("content : ", topicMessageDto.Content)
	fmt.Println("sender : ", topicMessageDto.Sender)
	fmt.Println("traceId : ", topicMessageDto.TraceID)
	fmt.Println("userId : ", topicMessageDto.UserID)
	for i, v := range topicMessageDto.Raws {
		fmt.Println("receiver", i, " : ", v)
	}
	fmt.Println("=====================================================")

	return sendMessageDataList, nil

}
