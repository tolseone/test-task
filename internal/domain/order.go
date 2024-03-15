package domain

import "time"

type Order struct {
	HotelID   string    `json:"hotel_id" validate:"required"`
	RoomID    string    `json:"room_id" validate:"required"`
	UserEmail string    `json:"email" validate:"required,email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}
