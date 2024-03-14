package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"applicationDesignTest/internal/domain/models"
	"applicationDesignTest/internal/service"

)

const (
	createOrderURL = "/orders"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) RegisterRoutes(r chi.Router) {
	r.Post(createOrderURL, h.CreateOrder)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order

	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.orderService.CreateOrder(newOrder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}
