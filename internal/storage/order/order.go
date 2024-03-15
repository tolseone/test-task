package order

import model "applicationDesignTest/internal/domain/order"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

var orders []model.Order

func (s *Storage) GetOrders() []model.Order {
	return orders
}

func (s *Storage) SaveOrder(order model.Order) bool {
	orders = append(orders, order)
	return true
}
