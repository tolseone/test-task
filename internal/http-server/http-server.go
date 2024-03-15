package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	model "applicationDesignTest/internal/domain/order"
	"applicationDesignTest/internal/lib/logger"

)

type Order interface {
	CreateOrder(order model.Order) (err error)
}

type ServerAPI struct {
	order     Order
	validator *validator.Validate
	logger    *logger.Logger
}

func Register(order Order) *ServerAPI {
	return &ServerAPI{
		order:     order,
		validator: validator.New(),
		logger:    logger.New(),
	}
}

func (s *ServerAPI) CreateOrder(w http.ResponseWriter, r *http.Request) {
	const op = "httpserver.CreateOrder"

	var newOrder model.Order

	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		s.logger.LogErrorf("%s: %s", op, err.Error())
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := s.validator.Struct(newOrder); err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			s.logger.LogErrorf("Validation error: %s", e)
		}
		http.Error(w, "Validation Error", http.StatusBadRequest)
		return
	}

	if err := s.order.CreateOrder(newOrder); err != nil {
		s.logger.LogErrorf("%s: %s", op, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.logger.LogInfo("Order created successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}
