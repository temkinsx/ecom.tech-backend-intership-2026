package main

import (
	"log"
	"net/http"
	"time"

	"ecom.tech-backend-intership-2026/internal/repository"
	"ecom.tech-backend-intership-2026/internal/service"
	"ecom.tech-backend-intership-2026/internal/transport/httpapi"
)

func main() {
	repo := repository.NewTodoRepository()
	svc := service.NewTodoService(repo)
	handler := httpapi.NewTodoHandler(svc)
	router := httpapi.NewRouter(handler)
	router = httpapi.Logging(router)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
