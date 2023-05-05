package service

import (
	"encoding/json"
	"log"

	"example.com/test_axxonsoft/v2/dto"
	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

func GetRabbitConnection() (*amqp.Connection, error) {
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
	TaskService *TaskService
}

func InitRabbitContext() *RabbitContext {
	var rabbitContext = new(RabbitContext)
	rabbitContext.initTaskConsumer()
	rabbitContext.initPublisher()
	return rabbitContext
}

func (r *RabbitContext) TaskHandler(taskDto dto.TaskDTO) error {
	return r.TaskService.ReceiveFromQueue(taskDto)
}

func (r *RabbitContext) initPublisher() {
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
		rabbitmq.WithPublisherOptionsExchangeName("task_exchange"),
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
			var taskDto = dto.TaskDTO{}
			var taskJson = string(d.Body)
			json.Unmarshal([]byte(taskJson), &taskDto)
			err := r.TaskHandler(taskDto)
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
