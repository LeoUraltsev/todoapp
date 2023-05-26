package handler

import (
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/handlers"
)

type Handler struct {
	//TODO: Logger
}

func New() handlers.Handler {
	return &Handler{}
}

func (h *Handler) Register() {
	http.HandleFunc("/tasks", h.Tasks)
}

func (h *Handler) Tasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello !"))
}