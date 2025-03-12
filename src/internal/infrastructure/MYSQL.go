package infrastructure

import (
    "database/sql"
    "pedidos/src/internal/domain/entities"
    "pedidos/src/internal/domain"
)

type messageRepository struct {
    db *sql.DB
}

func NewMessageRepository(db *sql.DB) domain.MessageRepository {
    return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(content string) (entities.Message, error) {
    result, err := r.db.Exec("INSERT INTO messages (content) VALUES (?)", content)
    if err != nil {
        return entities.Message{}, err
    }

    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return entities.Message{}, err
    }

    message := entities.Message{ID: int(lastInsertID), Content: content}
    return message, nil
}

func (r *messageRepository) GetLastMessage() (entities.Message, error) {
    var message entities.Message
    err := r.db.QueryRow("SELECT id, content FROM messages ORDER BY id DESC LIMIT 1").Scan(&message.ID, &message.Content)
    if err != nil {
        return entities.Message{}, err
    }
    return message, nil
}