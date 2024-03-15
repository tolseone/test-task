package order

import "applicationDesignTest/internal/domain"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

var orders []domain.Order

func (s *Storage) GetOrders() []domain.Order {
	return orders
}

func (s *Storage) SaveOrder(order domain.Order) bool {
	orders = append(orders, order)
	return true
}
