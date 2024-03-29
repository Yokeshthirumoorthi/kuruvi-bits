package main

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
	"encoding/json"
	"github.com/kuruvi-bits/transform/services"
	"github.com/kuruvi-bits/transform/utils"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func workflow(message utils.Message) {
	exif := services.Exif(message)
	fmt.Printf("Exif: %+v", exif)
	services.Resize(message)
	faces := services.DetectFaces(message)
	for index, boundingBox := range faces.Boxes {
		services.CropFace(message, boundingBox, index)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var message utils.Message
			err := json.Unmarshal([]byte(d.Body), &message)
			fmt.Println(err)
			log.Printf("Received a message: %s", d.Body)
			fmt.Printf("Received a message struct: %+v", message)
			workflow(message)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
