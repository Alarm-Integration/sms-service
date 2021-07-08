package controller

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Consumer(t *testing.T) {

	Convey("Given empty kafka configuration", t, func() {
		config := &kafka.ConfigMap{}

		Convey("When trying to connect with kafka server", func() {
			consumer, err := createConsumer(config)

			Convey("Then connection would be failed", func() {
				assert.Nil(t, consumer)
				assert.IsType(t, &kafka.Consumer{}, consumer)
				assert.EqualError(t, err, createClientErr.Error())
			})
		})
	})
}

//func Test_Subscribe_Topics(t *testing.T) {
//
//	Convey("Given no groupId to kafka server", t, func() {
//		consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
//			"bootstrap.servers": "http://kafka-server:9092/",
//			"auto.offset.reset": "earliest",
//		})
//		topics := []string{"test01", "test02"}
//
//		Convey("When trying to subscribing topics", func() {
//			err := subscribeTopics(consumer, topics)
//
//			Convey("Then subscribing topics would be failed", func() {
//				assert.EqualError(t, err, subscribeTopicErr.Error())
//			})
//		})
//	})
//}

//func Test_Kafka_Consumer_Subscribe_Faisl(t *testing.T) {
//
//	Convey("Given a groupId with empty string to kafka server", t, func() {
//		kafkaServer := "139.150.75.240"
//		//groupId := ""
//		topics := []string{"sms"}
//
//		Convey("When trying to connect with kafka server", func() {
//			err := ConnectKafkaConsumer(&kafka.ConfigMap{
//				"bootstrap.servers": kafkaServer,
//				//"group.id":          groupId,
//				"auto.offset.reset": "earliest",
//			}, topics, true)
//
//			Convey("Then subscribing topic message would be failed", func() {
//				assert.NotNil(t, err)
//				assert.EqualError(t, err, "[Kafka] consumer topic subscribe failed")
//			})
//		})
//	})
//}
