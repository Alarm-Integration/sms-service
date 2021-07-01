package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func StartListener(amqpURI, exchange, exchangeType, queueName, bindingKey, consumerTag string) {
	err := newConsumer(amqpURI, exchange, exchangeType, queueName, bindingKey, consumerTag)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("running forever")
	select {}

}

type consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

func newConsumer(amqpURI, exchange, exchangeType, queueName, bindingKey, consumerTag string) error {
	c := &consumer{
		conn:    nil,
		channel: nil,
		tag:     consumerTag,
		done:    make(chan error),
	}

	var err error

	log.Printf("dialing %s", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%s)", exchange)
	if err = c.channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %s", queueName)
	queue, err := c.channel.QueueDeclare(
		queueName, // name of the queue
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("queue Declare: %s", err)
	}

	log.Printf("declared Queue (%s %d messages, %d consumers), binding to Exchange (key %s)",
		queue.Name, queue.Messages, queue.Consumers, bindingKey)

	if err = c.channel.QueueBind(
		queue.Name, // name of the queue
		bindingKey, // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return fmt.Errorf("queue Bind: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %s)", c.tag)
	deliveries, err := c.channel.Consume(
		queue.Name, // name
		c.tag,      // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("queue Consume: %s", err)
	}

	go handle(deliveries, c.done)

	return nil
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"got %dB consumerTag: [%v] deliveryTag: [%v] routingkey: [%v] %s",
			len(d.Body),
			d.ConsumerTag,
			d.DeliveryTag,
			d.RoutingKey,
			d.Body,
		)
		handleRefreshEvent(d.Body, d.ConsumerTag)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}

func handleRefreshEvent(body []byte, consumerTag string) {
	updateToken := &UpdateToken{}
	err := json.Unmarshal(body, updateToken)
	if err != nil {
		log.Printf("Problem parsing UpdateToken: %v", err.Error())
	} else {
		log.Println("Reloading Viper config from Spring Cloud Config server")

		// Consumertag is same as application name.
		LoadConfigurationFromBranch(
			viper.GetString("configServerUrl"),
			consumerTag,
			viper.GetString("profile"),
			viper.GetString("configBranch"))
	}
}

type UpdateToken struct {
	Type               string `json:"type"`
	Timestamp          int    `json:"timestamp"`
	OriginService      string `json:"originService"`
	DestinationService string `json:"destinationService"`
	Id                 string `json:"id"`
}
