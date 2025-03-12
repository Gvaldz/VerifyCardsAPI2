package application

import (
    "pedidos/src/internal/domain/entities"
    "pedidos/src/internal/domain"
)

type CreateMessageUseCase struct {
    repo domain.MessageRepository
}

func NewCreateMessageUseCase(repo domain.MessageRepository) *CreateMessageUseCase {
    return &CreateMessageUseCase{repo: repo}
}

func (uc *CreateMessageUseCase) Execute(content string) (entities.Message, error) {
    return uc.repo.CreateMessage(content)
}