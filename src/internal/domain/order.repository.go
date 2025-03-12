package domain

import "pedidos/src/internal/domain/entities"

type MessageRepository interface {
    CreateMessage(content string) (entities.Message, error)
    GetLastMessage() (entities.Message, error) 
}