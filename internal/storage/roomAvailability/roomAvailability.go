package roomavailability

import model "applicationDesignTest/internal/domain/room_availability"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

var roomAvailabilities []model.RoomAvailability{
    {
        HotelID: "reddison", 
        RoomID: "lux", 
        Date: helper.Date(2024, 1, 1), 
        Quota: 1,
    },
    {
        HotelID: "reddison", 
        RoomID: "lux", 
        Date: helper.Date(2024, 1, 2), 
        Quota: 1,
    },
    {
        HotelID: "reddison", 
        RoomID: "lux",
        Date: helper.Date(2024, 1, 3), 
        Quota: 1,
    },
    {
        HotelID: "reddison", 
        RoomID: "lux", 
        Date: helper.Date(2024, 1, 4), 
        Quota: 1,
    },
    {
        HotelID: "reddison", 
        RoomID: "lux", 
        Date: helper.Date(2024, 1, 5), 
        Quota: 0,
    },
}

func (s *Storage) GetRoomAvailabilities() []model.RoomAvailability {
	
}

