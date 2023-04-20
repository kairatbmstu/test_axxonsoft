package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

func GetConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn, err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitContext struct {
	Publisher   *rabbitmq.Publisher
	Consumer    *rabbitmq.Consumer
	TaskHandler MessageHandler
}

func InitRabbitContext() *RabbitContext {
	var rabbitContext = new(RabbitContext)
	rabbitContext.initTaskConsumer()
	rabbitContext.initPublisher()
	return rabbitContext
}
