package infrastructure

import (
    "encoding/json"
    "net/http"
)

type MessageController struct {
    consumer *RabbitMQConsumer
}

func NewMessageController(consumer *RabbitMQConsumer) *MessageController {
    return &MessageController{consumer: consumer}
}

func (c *MessageController) GetMessages(w http.ResponseWriter, r *http.Request) {
    messages := c.consumer.GetMessages()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(messages)
}