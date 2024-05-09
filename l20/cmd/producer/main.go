package main

import (
	"context"
	"encoding/json"
	"log"
	"rabbit/internal/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewProducer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"l20.fruits", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		name, size := utils.GetRandomFruit()
		fruit := utils.Fruit{
			Name: name,
			Size: size,
		}

		body, err := json.Marshal(fruit)
		if err != nil {
			utils.FailOnError(err, "Failed to marshal fruit")
		}

		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		utils.FailOnError(err, "Failed to publish a message")
		log.Printf("Sent %s\n", body)

		cancel()

		time.Sleep(time.Second)
	}
}

func main() {
	NewProducer()
}
