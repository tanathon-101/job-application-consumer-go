package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type JobApplication struct {
	ApplicantName string `json:"applicantName"`
	Position      string `json:"position"`
	Email         string `json:"email"`
}

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"job_application", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer tag
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var app JobApplication
			if err := json.Unmarshal(d.Body, &app); err != nil {
				log.Println("Error decoding message:", err)
				continue
			}
			fmt.Printf("📩 New Application: %+v\n", app)
		}
	}()

	log.Println("🟢 Waiting for messages. To exit press CTRL+C")
	<-forever
}
