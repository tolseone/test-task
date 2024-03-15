// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"errors"
	"net/http"
	"os"

	"applicationDesignTest/internal/domain/models"
	"applicationDesignTest/internal/lib/helper"
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/router"
)

var Orders = []models.Order{}

var Availability = []models.RoomAvailability{
	{HotelID: "reddison", RoomID: "lux", Date: helper.Date(2024, 1, 1), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: helper.Date(2024, 1, 2), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: helper.Date(2024, 1, 3), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: helper.Date(2024, 1, 4), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: helper.Date(2024, 1, 5), Quota: 0},
}

func main() {
	// Get logger
	logger := logger.New()
	logger.LogInfo("logger initialized")

	// Get router with registred routes
	router := router.GetRouter()
	logger.LogInfo("router initialized")

	logger.LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if errors.Is(err, http.ErrServerClosed) {
		logger.LogInfo("Server closed")
	} else if err != nil {
		logger.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
