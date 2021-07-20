package controller

import (
	"errors"
	"fmt"

	smsService "github.com/GreatLaboratory/go-sms/service"
	"github.com/GreatLaboratory/go-sms/util"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	createClientErr   = errors.New("[Kafka] consumer create failed")
	subscribeTopicErr = errors.New("[Kafka] consumer topic subscribe failed")
)

func ConnectKafkaConsumer(config *kafka.ConfigMap, topics []string) error {
	consumer, createErr := createConsumer(config)
	if createErr != nil {
		return createErr
	}

	subscribeErr := subscribeTopics(consumer, topics)
	if subscribeErr != nil {
		return subscribeErr
	}

	fmt.Println("[Kafka] Connection Success")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			// The client will automatically try to recover from all errors.
			fmt.Printf("[Kafka] Connection Error: %v (%v)\n", err, msg)
			continue
		}

		fmt.Println("[Kafka] Consumed Message Topic Partition : ", msg.TopicPartition)
		fmt.Println("[Kafka] Consumed Message Topic Value : ", string(msg.Value))
		requestBody, requestID, err := util.ConvertByteToDtoList(msg.Value)
		if err != nil {
			fmt.Println("[Kafka] Convert Error : ", err)
			continue
		}
		for _, sendMessageDto := range requestBody.Messages {
			go smsService.SendMessage(sendMessageDto, requestID)
		}
		if err != nil {
			fmt.Println("[SMS] Send Error : ", err)
		}
	}

}

func createConsumer(config *kafka.ConfigMap) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		fmt.Println(err)
		return nil, createClientErr
	}
	return consumer, nil
}

func subscribeTopics(consumer *kafka.Consumer, topics []string) error {
	err := consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return subscribeTopicErr
	}
	return nil
}
