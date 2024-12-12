package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/interface/controller"
	"github.com/swaggo/http-swagger" // Swagger dokümantasyonu için
)

// @title Go Clean Architecture API
// @version 1.0
// @description This is a sample server for demonstrating Clean Architecture with Go.
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	// API endpoint'leri burada eklenecek
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controller.GetUserByID).Methods("GET")

	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
