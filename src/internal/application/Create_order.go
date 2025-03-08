package application

import (
    "pedidos/src/internal/domain/entities"
    "pedidos/src/internal/domain"
)

type CreateOrderUseCase struct {
    repo domain.OrderRepository
}

func NewCreateOrderUseCase(repo domain.OrderRepository) *CreateOrderUseCase {
    return &CreateOrderUseCase{repo: repo}
}

func (uc *CreateOrderUseCase) Execute() (entities.Order, error) {
    return uc.repo.CreateOrder()
}