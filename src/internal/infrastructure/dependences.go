package infrastructure

import (
    "database/sql"
    "pedidos/src/core"
)

type Dependencies struct {
    DB      *sql.DB
    RabbitMQ *core.RabbitMQConnection
}

func NewDependencies(db *sql.DB, rabbitMQ *core.RabbitMQConnection) *Dependencies {
    return &Dependencies{DB: db, RabbitMQ: rabbitMQ}
}

func (d *Dependencies) GetConsumer() *RabbitMQConsumer {
    return NewRabbitMQConsumer(d.RabbitMQ)
}

func (d *Dependencies) GetMessageRoutes() *MessageRoutes {
    consumer := NewRabbitMQConsumer(d.RabbitMQ)
    messageController := NewMessageController(consumer)
    return NewMessageRoutes(messageController)
}