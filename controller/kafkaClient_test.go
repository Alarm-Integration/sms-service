package controller

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Connect_Kafka_Consumer_Fail_By_Create_Error(t *testing.T) {

	Convey("Given empty kafka configuration", t, func() {
		config := &kafka.ConfigMap{}
		topics := []string{"test01", "test02"}

		Convey("When trying to connect with kafka server", func() {
			err := ConnectKafkaConsumer(config, topics)

			Convey("Then connection would be failed", func() {
				assert.EqualError(t, err, createClientErr.Error())
			})
		})
	})
}

func Test_Connect_Kafka_Consumer_Fail_By_Subscribe_Error(t *testing.T) {

	Convey("Given no groupId to kafka server", t, func() {
		config := &kafka.ConfigMap{
			"bootstrap.servers": "139.150.75.239:19092/",
			"group.id":          "",
			"auto.offset.reset": "earliest",
		}
		topics := []string{"test01", "test02"}

		Convey("When trying to subscribing topics", func() {
			err := ConnectKafkaConsumer(config, topics)

			Convey("Then subscribing topics would be failed", func() {
				assert.EqualError(t, err, subscribeTopicErr.Error())
			})
		})
	})
}
