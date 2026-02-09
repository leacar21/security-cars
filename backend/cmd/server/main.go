package main

import (
	"log"
	"net/http"

	"github.com/example/security-cars/backend/internal/adapters/http"
	"github.com/example/security-cars/backend/internal/application/usecase"
)

func main() {
	greetingService := usecase.NewGreetingService()
	handler := httpadapter.NewHandler(greetingService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.Routes(),
	}

	log.Println("server listening on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
