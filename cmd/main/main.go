package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeoUraltsev/todoapp/internal/task/handler"
)

func main() {
	
	fmt.Println("app is running ...")

	h := handler.New()
	h.Register()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
