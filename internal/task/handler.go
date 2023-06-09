package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/handlers"
	"github.com/LeoUraltsev/todoapp/pkg/logger"
)

type Handler struct {
	logger     *logger.Logger
	repository Repository
}

func NewHandler(logger *logger.Logger, repository Repository) handlers.Handler {
	return &Handler{
		logger:     logger,
		repository: repository,
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.GET("/tasks", h.GetTask)
	router.GET("/tasks/:id", h.GetTaskByID)
	router.POST("/tasks", h.CreateTask)
	router.PUT("/tasks/:id", h.FullUpdateTask)
	router.PATCH("/tasks/:id", h.PartiallyUpdateTask)
	router.DELETE("/tasks/:id", h.DeleteTask)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Debug("GET /tasks")
	w.Header().Set("Content-Type", "application/json")

	tasks, err := h.repository.FindAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	resp, err := json.Marshal(tasks)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Debug("POST /tasks")

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var task Task
	err = json.Unmarshal(resp, &task)
	if err != nil {
		h.logger.Sugar().Errorf("unmarshal err %v", err)
		return
	}

	h.logger.Debug(fmt.Sprintf("%s", task))

	id, err := h.repository.Create(context.Background(), task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Sugar().Errorf("create task error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	task, err := h.repository.FindOne(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		h.logger.Sugar().Infof("task by id: %s not found. err: %v", id, err)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Sugar().Infof("marshal error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func (h *Handler) FullUpdateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "Application-json")

	var task Task

	task.ID = params.ByName("id")

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = json.Unmarshal(resp, &task)
	if err != nil {
		h.logger.Sugar().Errorf("unmarshal err %v", err)
		return
	}

	h.logger.Debug(fmt.Sprintf("%s", task))

	err = h.repository.Update(context.Background(), task)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}

func (h *Handler) PartiallyUpdateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "Application-json")

	var task Task

	task.ID = params.ByName("id")

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = json.Unmarshal(resp, &task)
	if err != nil {
		h.logger.Sugar().Errorf("unmarshal err %v", err)
		return
	}

	h.logger.Debug(fmt.Sprintf("%s", task))

	err = h.repository.Update(context.Background(), task)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	err := h.repository.Delete(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

}
