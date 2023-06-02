package main

import (
	"context"
	"fmt"
	"github.com/LeoUraltsev/todoapp/internal/task"
	"github.com/LeoUraltsev/todoapp/internal/task/db"
	"github.com/LeoUraltsev/todoapp/pkg/client/postgesql"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/config"
	"github.com/LeoUraltsev/todoapp/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	log.Info("app is running ...")

	log.Info("read config ...")
	cfg := config.GetInstance()

	log.Info("connect to database ...")

	pgCfg := postgesql.Config{
		Username: cfg.Username,
		Password: cfg.Password,
		Host:     cfg.Host,
		Port:     cfg.Port,
		Db:       cfg.Db,
	}

	client, err := postgesql.NewClient(context.Background(), pgCfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.Background())
	if err != nil {
		log.Error(fmt.Sprintf("error is ping to database: %v", err))
	}

	repository := db.NewRepository(client, log)

	r := httprouter.New()

	log.Info("creating handler ...")
	h := task.NewHandler(log, repository)

	log.Info("routes register ...")
	h.Register(r)

	start(log, cfg, r)

}

func start(logger *logger.Logger, cfg *config.Config, router *httprouter.Router) {
	logger.Info(fmt.Sprintf("server is start on host: %s and port: %s", cfg.Listen.Host, cfg.Listen.Port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), router); err != nil {
		logger.Fatal(err.Error())
	}
}
