package infrastructure

import (
	"log"
	"sync"
	"pedidos/src/core"
	"pedidos/src/internal/application"
	"pedidos/src/internal/domain/entities"
)

type RabbitMQConsumer struct {
	conn                *core.RabbitMQConnection
	createMessageUseCase *application.CreateMessageUseCase
	producer            *RabbitMQProducer
	messages            []entities.Message
	mutex               sync.Mutex
}

func NewRabbitMQConsumer(conn *core.RabbitMQConnection, createMessageUseCase *application.CreateMessageUseCase, producer *RabbitMQProducer) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		conn:                conn,
		createMessageUseCase: createMessageUseCase,
		producer:            producer,
		messages:            make([]entities.Message, 0),
	}
}

func (c *RabbitMQConsumer) Start() {
	log.Println("Connecting to RabbitMQ...")

	inputQueue, err := c.conn.Ch.QueueDeclare(
		"cards", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare input queue: %s", err)
	}

	log.Printf("Queue '%s' declared successfully", inputQueue.Name)

	msgs, err := c.conn.Ch.Consume(
		inputQueue.Name, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	log.Println("Waiting for messages...")

	for d := range msgs {
		content := string(d.Body)
		message, err := c.createMessageUseCase.Execute(content)
		if err != nil {
			log.Printf("Failed to save message: %s", err)
			continue
		}

		c.mutex.Lock()
		c.messages = append(c.messages, message)
		c.mutex.Unlock()
		log.Printf("Received and saved message: %s", message.Content)

		newMessageContent := "Processed: " + message.Content

		err = c.producer.PublishMessage("messages", newMessageContent)
		if err != nil {
			log.Printf("Failed to publish processed message: %s", err)
			continue
		}

		log.Printf("Published processed message to queue: %s", newMessageContent)
	}
}