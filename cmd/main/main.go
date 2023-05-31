package main

import (
	"context"
	"fmt"
	"github.com/LeoUraltsev/todoapp/internal/task/model"
	"github.com/LeoUraltsev/todoapp/pkg/client/postgesql"
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/config"
	"github.com/LeoUraltsev/todoapp/internal/task/handler"
	"github.com/LeoUraltsev/todoapp/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	log.Info("app is running ...")

	log.Info("read config ...")
	cfg := config.GetInstance()

	log.Info("connect to database ...")
	db, err := postgesql.NewClient(context.Background(), cfg.StorageConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Error(fmt.Sprintf("error is ping to database: %v", err))
	}
	var task model.Task

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info(task.Description)

	log.Info("creating handler ...")
	h := handler.New(log)

	log.Info("routes register ...")
	h.Register()

	start(log, cfg)

}

func start(logger *logger.Logger, cfg *config.Config) {
	logger.Info(fmt.Sprintf("server is start on host: %s and port: %s", cfg.Listen.Host, cfg.Listen.Port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), nil); err != nil {
		logger.Fatal(err.Error())
	}
}
