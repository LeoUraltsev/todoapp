package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/config"
	"github.com/LeoUraltsev/todoapp/internal/task/handler"
)

func main() {

	fmt.Println("app is running ...")

	fmt.Println("read config ...")
	cfg := config.GetInstance()

	h := handler.New()
	h.Register()

	fmt.Printf("server is start on host: %s and port: %s", cfg.Listen.Host, cfg.Listen.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), nil))
}
