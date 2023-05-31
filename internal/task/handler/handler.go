package handler

import (
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/handlers"
	"github.com/LeoUraltsev/todoapp/pkg/logger"
)

type Handler struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) handlers.Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Register() {
	http.HandleFunc("/tasks", h.Tasks)
}

func (h *Handler) Tasks(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("Accessing the route /tasks")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello !"))
}