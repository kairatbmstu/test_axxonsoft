package rabbit

import (
	"log"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)


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
