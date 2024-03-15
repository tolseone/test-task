package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	model "applicationDesignTest/internal/domain/order"
	resp "applicationDesignTest/internal/lib/api/response"
	"applicationDesignTest/internal/lib/logger"

)

type Request struct {
	HotelID   string    `json:"hotel_id" validate:"required"`
	RoomID    string    `json:"room_id" validate:"required"`
	UserEmail string    `json:"email" validate:"required,email"`
	From      time.Time `json:"from" validate:"required,time"`
	To        time.Time `json:"to" validate:"required,time"`
}

type Response struct {
	resp.Response
}

type Order struct {
	orders       []model.Order
	availability map[string]map[string]map[time.Time]int
}

type RepositoryOrder interface { // OrderCreator
	CreateOrder(order model.Order) (err error)
}

func New1(log *logger.Logger, repoOrder RepositoryOrder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "service.order.New"

		var req Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.LogErrorf("failed to decode request body: %s", err.Error())

			
		}
		

	}
}

func New() *Order {
	return &Order{
		orders:       []model.Order{},
		availability: make(map[string]map[string]map[time.Time]int),
	}
}

func (ordr *Order) CreateOrder(newOrder model.Order) error {
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
