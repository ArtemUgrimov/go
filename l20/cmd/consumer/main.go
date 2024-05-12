package main

import (
	"encoding/json"
	"log"
	"rabbit/internal/utils"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"l20.fruits", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	var lock sync.RWMutex = sync.RWMutex{}
	var fruits map[string]map[string]int32 = make(map[string]map[string]int32)

	go func() {
		for d := range msgs {
			// log.Printf("Received a message: %s", d.Body)

			fruit := utils.Fruit{}
			err := json.Unmarshal(d.Body, &fruit)
			utils.FailOnError(err, "Failed to unmarshal fruit")

			category := utils.SMALL
			if fruit.Size > 33 && fruit.Size < 67 {
				category = utils.MEDIUM
			} else if fruit.Size >= 67 {
				category = utils.LARGE
			}

			lock.Lock()
			_, ok := fruits[fruit.Name]
			if !ok {
				fruits[fruit.Name] = make(map[string]int32)
			}

			fruits[fruit.Name][category] += 1
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.RLock()
			for k, v := range fruits {
				log.Printf("%s: %s:%d, %s: %d, %s: %d\n", k, utils.SMALL, v[utils.SMALL], utils.MEDIUM, v[utils.MEDIUM], utils.LARGE, v[utils.LARGE])
			}
			log.Printf("==============\n\n")
			lock.RUnlock()
			time.Sleep(time.Second * 10)
		}
	}()

	<-forever
}

func main() {
	NewConsumer()
}
