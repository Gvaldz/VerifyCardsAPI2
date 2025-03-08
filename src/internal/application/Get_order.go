package application

import (
    "pedidos/src/internal/domain/entities"
    "pedidos/src/internal/domain"
)

type GetOrderUseCase struct {
    repo domain.OrderRepository
}

func NewGetOrderUseCase(repo domain.OrderRepository) *GetOrderUseCase {
    return &GetOrderUseCase{repo: repo}
}

func (uc *GetOrderUseCase) Execute(orderID int) (entities.Order, error) {
    return uc.repo.GetOrderByID(orderID)
}