package router

import (
	"github.com/go-chi/chi"

	"applicationDesignTest/internal/http-server"
	"applicationDesignTest/internal/services"

)

const (
	createOrderURL = "/orders"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()

	orderHandler := httpserver.Register(services.NewOrderService())

	router.Post(createOrderURL, orderHandler.CreateOrder)

	return router
}
