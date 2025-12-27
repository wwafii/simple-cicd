package main

import (
	"log"
	"net/http"
	"simple-cicd/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.HealthCheck)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	log.Println("User API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
