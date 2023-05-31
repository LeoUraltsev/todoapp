package main

import (
	"fmt"
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/config"
	"github.com/LeoUraltsev/todoapp/internal/task/handler"
	"github.com/LeoUraltsev/todoapp/pkg/logger"
)

func main() {
	logger := logger.GetLogger()

	logger.Info("app is running ...")

	logger.Info("read config ...")
	cfg := config.GetInstance()

	logger.Info("creating handler ...")
	h := handler.New(logger)

	logger.Info("routes register ...")
	h.Register()

	start(logger, cfg)

}

func start(logger *logger.Logger, cfg *config.Config) {
	logger.Info(fmt.Sprintf("server is start on host: %s and port: %s", cfg.Listen.Host, cfg.Listen.Port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), nil); err != nil {
		logger.Fatal(err.Error())
	}
}
