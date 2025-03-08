package domain

import "pedidos/src/internal/domain/entities"

type OrderRepository interface {
    CreateOrder() (entities.Order, error)
    GetOrderByID(orderID int) (entities.Order, error)
}