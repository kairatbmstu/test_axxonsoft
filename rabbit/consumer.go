package rabbit

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

var consumer *rabbitmq.Consumer

type MessageHandler func(message string) error

func InitConsumer(messageHandler MessageHandler) {
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err = rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			msg := string(d.Body)
			err := messageHandler(msg)
			if err != nil {
				return rabbitmq.NackDiscard
			}
			return rabbitmq.Ack
		},
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("my_routing_key"),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")

}
