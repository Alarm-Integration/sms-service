package controller

import (
	"errors"
	"fmt"

	smsService "github.com/GreatLaboratory/go-sms/service"
	"github.com/GreatLaboratory/go-sms/util"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConnectKafkaConsumer(kafkaServer, groupId string, topics []string, isTest ...bool) error {
	consumer, createErr := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if createErr != nil {
		return errors.New("consumer create failed")
	}

	subscribeErr := consumer.SubscribeTopics(topics, nil)
	if subscribeErr != nil {
		return errors.New("consumer topic subscribe failed")
	}

	fmt.Println("Kafka Connection Success")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			if len(isTest) != 0 && isTest[0] {
				consumer.Close()
				return errors.New(err.Error())
			}
			// The client will automatically try to recover from all errors.
			fmt.Printf("Kafka Connection Error: %v (%v)\n", err, msg)
		} else {
			fmt.Println("Consumed Message Topic Partition : ", msg.TopicPartition)
			fmt.Println("Consumed Message Topic Value : ", string(msg.Value))

			params := make(map[string]string)
			sendMessageDataList, err := util.ConvertByteToDtoList(msg.Value)
			if err != nil {
				fmt.Println("Convert Error : ", err)
			} else {
				err := smsService.SendGroupMessage(params, sendMessageDataList)
				if err != nil {
					fmt.Println("Send SMS Error : ", err)
				}
			}
		}
	}
}
