package infrastructure

import (
	"database/sql"
	"pedidos/src/core"
	"pedidos/src/internal/application"
)

type Dependencies struct {
	DB                 *sql.DB
	RabbitMQ           *core.RabbitMQConnection
	createMessageUseCase *application.CreateMessageUseCase
	getMessageUseCase  *application.GetMessageUseCase
	producer           *RabbitMQProducer
}

func NewDependencies(db *sql.DB, rabbitMQ *core.RabbitMQConnection) *Dependencies {
	repo := NewMessageRepository(db)

	createMessageUseCase := application.NewCreateMessageUseCase(repo)
	getMessageUseCase := application.NewGetMessageUseCase(repo)

	producer := NewRabbitMQProducer(rabbitMQ)

	return &Dependencies{
		DB:                 db,
		RabbitMQ:           rabbitMQ,
		createMessageUseCase: createMessageUseCase,
		getMessageUseCase:  getMessageUseCase,
		producer:           producer,
	}
}

func (d *Dependencies) GetMessageRoutes() *MessageRoutes {
	messageController := NewMessageController(d.getMessageUseCase, d.createMessageUseCase)
	return NewMessageRoutes(messageController)
}

func (d *Dependencies) GetMessageConsumer() *RabbitMQConsumer {
	return NewRabbitMQConsumer(d.RabbitMQ, d.createMessageUseCase, d.producer)
}