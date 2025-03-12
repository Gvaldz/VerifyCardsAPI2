package infrastructure

import (
	"log"
	"pedidos/src/core"
	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	conn *core.RabbitMQConnection
}

func NewRabbitMQProducer(conn *core.RabbitMQConnection) *RabbitMQProducer {
	return &RabbitMQProducer{
		conn: conn,
	}
}

func (p *RabbitMQProducer) PublishMessage(queueName string, message string) error {
	err := p.conn.Ch.Publish(
		"",         // exchange
		queueName,  // routing key (queue name)
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Failed to publish message to queue '%s': %s", queueName, err)
		return err
	}

	log.Printf("Published message to queue '%s': %s", queueName, message)
	return nil
}