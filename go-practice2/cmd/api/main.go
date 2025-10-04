package main

import (
	"log"
	"net/http"

	"practice2/internal/handlers"
	"practice2/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Регистрируем обработчик для /user и оборачиваем его middleware
	mux.Handle("/user", middleware.Auth(http.HandlerFunc(handlers.UserHandler)))

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
