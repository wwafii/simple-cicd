package handlers

import (
    "encoding/json"
    "net/http"
    "simple-cicd/models"
)

var users = []models.User{
    {ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password123"},
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "OK", "service": "User API"})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var req models.RegisterRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    
    if req.Name == "" || req.Email == "" || req.Password == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }
    
    newUser := models.User{
        ID:       len(users) + 1,
        Name:     req.Name,
        Email:    req.Email,
        Password: req.Password, 
    }
    
    users = append(users, newUser)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var req models.LoginRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    for _, user := range users {
        if user.Email == req.Email && user.Password == req.Password {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]interface{}{
                "message": "Login successful",
                "user":    map[string]string{"name": user.Name, "email": user.Email},
            })
            return
        }
    }
    
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    
    user := users[0]
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
