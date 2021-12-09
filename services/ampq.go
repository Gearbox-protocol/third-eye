package services

import (
	"github.com/Gearbox-protocol/third-eye/config"
	logger "github.com/Gearbox-protocol/third-eye/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type ampqService struct {
	ch *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Service constructor
func NewAMQPService(config *config.Config) {
	if config.AMPQEnable == "0" {
		return
	}
	conn, err := amqp.Dial(config.AMPQUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	logger.SetAMQP(ch)
}
