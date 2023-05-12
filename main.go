package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("hello world")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ", err)
		log.Panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"testQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}

	fmt.Println(q)

	if err = ch.Publish(
		"",
		"testQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello Mars"),
		},
	) 

	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	fmt.Println("Successfully Published Message to Queue")
}
