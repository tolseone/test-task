package order

import (
	"errors"
	"fmt"
	"time"

	"applicationDesignTest/internal/domain/models"

)

type Order struct {
	orders       []models.Order
	availability map[string]map[string]map[time.Time]int
}

type RepositoryOrder interface {
	CreateOrder(order models.Order) (err error)
}

func New() *Order {
	return &Order{
		orders:       []models.Order{},
		availability: make(map[string]map[string]map[time.Time]int),
	}
}

func (ordr *Order) CreateOrder(newOrder models.Order) error {
	const op = "Service.CreateOrder"

	for d := newOrder.From; d.Before(newOrder.To); d = d.AddDate(0, 0, 1) {
		if quota, ok := ordr.availability[newOrder.HotelID][newOrder.RoomID][d]; !ok || quota < 1 {
			ordr.log.Error("hotel room is not available for selected dates")
			return fmt.Errorf("%s: %w", op, errors.New("hotel room is not available for selected dates"))
		}
	}

	for d := newOrder.From; d.Before(newOrder.To); d = d.AddDate(0, 0, 1) {
		ordr.availability[newOrder.HotelID][newOrder.RoomID][d] -= 1
	}

	ordr.orders = append(ordr.orders, newOrder)

	return nil
}
