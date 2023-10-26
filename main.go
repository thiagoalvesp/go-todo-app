package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/thiagoalvesp/go-todo-app/configs"
	"github.com/thiagoalvesp/go-todo-app/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(0)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
