package application

import (
    "pedidos/src/internal/domain/entities"
    "pedidos/src/internal/domain"
)

type GetMessageUseCase struct {
    repo domain.MessageRepository
}

func NewGetMessageUseCase(repo domain.MessageRepository) *GetMessageUseCase {
    return &GetMessageUseCase{repo: repo}
}

func (uc *GetMessageUseCase) Execute() (entities.Message, error) {
    return uc.repo.GetLastMessage()
}