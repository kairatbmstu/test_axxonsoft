package service

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

type MessageHandler func(message string) error

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


func (r RabbitContext) initPublisher() {
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	r.Publisher = publisher
}

func (r RabbitContext) ClosePublisher() {
	defer r.Publisher.Close()
}

func (r RabbitContext) SendTask(message string) error {
	err := r.Publisher.Publish(
		[]byte(message),
		[]string{"task_routing_key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("task_exchange"),
	)
	return err
}



func (r RabbitContext) initTaskConsumer() {
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			err := r.TaskHandler(string(d.Body))
			if err != nil {
				log.Printf("error happened while calling message handler: %v", err.Error())
				return rabbitmq.NackDiscard
			}
			return rabbitmq.Ack
		},
		"task_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("task_routing_key"),
		rabbitmq.WithConsumerOptionsExchangeName("task_exchange"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	r.Consumer = consumer
}
