package rabbit

import (
	"log"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

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
