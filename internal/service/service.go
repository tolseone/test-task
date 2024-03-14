package service

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"applicationDesignTest/internal/domain/models"
)

// OrderService представляет сервис для работы с заказами.
type OrderService struct {
	orders       []models.Order
	availability map[string]map[string]map[time.Time]int
	log          *slog.Logger
}

// NewOrderService создает новый экземпляр OrderService.
func NewOrderService(log *slog.Logger) *OrderService {
	return &OrderService{
		orders:       []models.Order{},
		availability: make(map[string]map[string]map[time.Time]int),
		log:          log,
	}
}

// CreateOrder создает новый заказ на бронирование номера.
func (s *OrderService) CreateOrder(newOrder models.Order) error {
	const op = "Service.CreateOrder"

	log := s.log.With(
		slog.String("op", op),
		slog.String("HotelID", newOrder.HotelID),
		slog.String("RoomID", newOrder.RoomID),
		slog.String("UserEmail", newOrder.UserEmail),
	)

	log.Info("attempting to create order")

	for d := newOrder.From; d.Before(newOrder.To); d = d.AddDate(0, 0, 1) {
		if quota, ok := s.availability[newOrder.HotelID][newOrder.RoomID][d]; !ok || quota < 1 {
			s.log.Error("hotel room is not available for selected dates")
			return fmt.Errorf("%s: %w", op, errors.New("hotel room is not available for selected dates"))
		}
	}

	for d := newOrder.From; d.Before(newOrder.To); d = d.AddDate(0, 0, 1) {
		s.availability[newOrder.HotelID][newOrder.RoomID][d] -= 1
	}

	s.orders = append(s.orders, newOrder)
	return nil
}
