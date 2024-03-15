package services

import (
	"errors"
	"fmt"

	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/storage/order"
	"applicationDesignTest/internal/storage/roomAvailability"

)

type RepositoryOrder interface {
	GetOrders() (orders []domain.Order)
	SaveOrder(order domain.Order) bool
}

type RepositoryRoomAvailability interface {
	GetRoomAvailabilities(hotelID string, roomID string) (roomAvailabilities []domain.RoomAvailability)
}

type OrderService struct {
	repoOrder            RepositoryOrder
	repoRoomAvailability RepositoryRoomAvailability
	logger               *logger.Logger
}

func NewOrderService() *OrderService {
	repoOrder := order.NewStorage()
	repoRoomAvailability := roomavailability.NewStorage()
	logger := logger.New()

	return &OrderService{
		repoOrder:            repoOrder,
		repoRoomAvailability: repoRoomAvailability,
		logger:               logger,
	}
}

func (ordr *OrderService) CreateOrder(newOrder domain.Order) error {
	const op = "Service.CreateOrder"

	for d := newOrder.From; d.Before(newOrder.To); d = d.AddDate(0, 0, 1) {
		availabilities := ordr.repoRoomAvailability.GetRoomAvailabilities(newOrder.HotelID, newOrder.RoomID)

		var quota int
		for _, availability := range availabilities {
			if availability.Date.Equal(d) {
				quota = availability.Quota
				break
			}
		}

		if quota < 1 {
			ordr.logger.LogErrorf("hotel room is not available for selected dates")
			return fmt.Errorf("%s: %w", op, errors.New("hotel room is not available for selected dates"))
		}
	}

	for d := newOrder.From; d.Before(newOrder.To); d = d.AddDate(0, 0, 1) {
		availabilities := ordr.repoRoomAvailability.GetRoomAvailabilities(newOrder.HotelID, newOrder.RoomID)

		for _, availability := range availabilities {
			if availability.Date.Equal(d) {
				availability.Quota--
				break
			}
		}

		// TODO: Сохранение обновленной доступности номеров в хранилище
	}

	if !ordr.repoOrder.SaveOrder(newOrder) {
		return fmt.Errorf("%s: %w", op, errors.New("failed to save order"))
	}

	return nil
}
