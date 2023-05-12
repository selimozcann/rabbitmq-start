package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	defer ch.Close()

	message, err := ch.Consume(
		"testQueue",
		"consumer",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range message {
			fmt.Println("Received Message: %s\n", string(d.Body))
		}
	}()

	fmt.Println("Succesfully Connected to RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
