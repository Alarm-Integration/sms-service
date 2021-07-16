package model

import (
	"github.com/streadway/amqp"
)

type UpdateToken struct {
	Type               string `json:"type"`
	Timestamp          int    `json:"timestamp"`
	OriginService      string `json:"originService"`
	DestinationService string `json:"destinationService"`
	Id                 string `json:"id"`
}

type Consumer struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Tag     string
	Done    chan error
}
