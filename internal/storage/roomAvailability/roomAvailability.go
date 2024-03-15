package roomavailability

import (
	"applicationDesignTest/internal/domain"
	"applicationDesignTest/internal/lib/helper"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

var roomAvailabilities = []domain.RoomAvailability{
	{
		HotelID: "reddison",
		RoomID:  "lux",
		Date:    helper.Date(2024, 1, 1),
		Quota:   1,
	},
	{
		HotelID: "reddison",
		RoomID:  "lux",
		Date:    helper.Date(2024, 1, 2),
		Quota:   1,
	},
	{
		HotelID: "reddison",
		RoomID:  "lux",
		Date:    helper.Date(2024, 1, 3),
		Quota:   1,
	},
	{
		HotelID: "reddison",
		RoomID:  "lux",
		Date:    helper.Date(2024, 1, 4),
		Quota:   1,
	},
	{
		HotelID: "reddison",
		RoomID:  "lux",
		Date:    helper.Date(2024, 1, 5),
		Quota:   0,
	},
}

func (s *Storage) GetRoomAvailabilities(hotelID string, roomID string) []domain.RoomAvailability {
	result := make([]domain.RoomAvailability, 0)

	for _, roomavailability := range roomAvailabilities {
		if hotelID == roomavailability.HotelID && roomID == roomavailability.RoomID {
			result = append(result, roomavailability)
		}
	}
	return result
}
