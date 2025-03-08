package infrastructure

import (
    "log"
    "sync"
    "pedidos/src/core"
    "pedidos/src/internal/domain/entities"
)

type RabbitMQConsumer struct {
    conn      *core.RabbitMQConnection
    messages  []entities.Message
    mutex     sync.Mutex
}

func NewRabbitMQConsumer(conn *core.RabbitMQConnection) *RabbitMQConsumer {
    return &RabbitMQConsumer{
        conn:     conn,
        messages: make([]entities.Message, 0),
    }
}

func (c *RabbitMQConsumer) Start() {
    log.Println("Connecting to RabbitMQ...")

    q, err := c.conn.Ch.QueueDeclare(
        "cards", // name
        true,   // durable
        false,  // delete when unused
        false,  // exclusive
        false,  // no-wait
        nil,    // arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %s", err)
    }

    log.Printf("Queue '%s' declared successfully", q.Name)

    msgs, err := c.conn.Ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %s", err)
    }

    log.Println("Waiting for messages...")

    for d := range msgs {
        message := entities.Message{Content: string(d.Body)}
        c.mutex.Lock()
        c.messages = append(c.messages, message)
        c.mutex.Unlock()
        log.Printf("Received message: %s", message.Content)
    }
}

func (c *RabbitMQConsumer) GetMessages() []entities.Message {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    return c.messages
}