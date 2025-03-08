package infrastructure

import (
    "database/sql"
    "pedidos/src/internal/domain/entities"
    "pedidos/src/internal/domain"
)

type orderRepository struct {
    db *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
    return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder() (entities.Order, error) {
    result, err := r.db.Exec("INSERT INTO orders DEFAULT VALUES")
    if err != nil {
        return entities.Order{}, err
    }

    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return entities.Order{}, err
    }

    order := entities.Order{ID: int(lastInsertID)}
    return order, nil
}

func (r *orderRepository) GetOrderByID(orderID int) (entities.Order, error) {
    var order entities.Order
    err := r.db.QueryRow("SELECT id FROM orders WHERE id = ?", orderID).Scan(&order.ID)
    if err != nil {
        return entities.Order{}, err
    }
    return order, nil
}