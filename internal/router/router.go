package router

import (
	"github.com/go-chi/chi"

	"applicationDesignTest/internal/http-server"
	"applicationDesignTest/internal/service/order"
)

const (
	createOrderURL = "/orders"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()

	orderHandler := httpserver.Register(order.New())

	router.Post(createOrderURL, orderHandler.CreateOrder)

	return router
}
