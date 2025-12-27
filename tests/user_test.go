package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "simple-cicd/handlers"
)

func TestHealthCheck(t *testing.T) {
    req, _ := http.NewRequest("GET", "/health", nil)
    rr := httptest.NewRecorder()
    handlers.HealthCheck(rr, req)
    
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected %d, got %d", http.StatusOK, status)
    }
    
    var result map[string]interface{}
    json.NewDecoder(rr.Body).Decode(&result)
    if result["status"] != "OK" {
        t.Errorf("expected status OK, got %v", result["status"])
    }
}

func TestRegisterSuccess(t *testing.T) {
    reqBody := `{"name":"Jane Doe","email":"jane@example.com","password":"secret"}`
    req, _ := http.NewRequest("POST", "/register", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    
    rr := httptest.NewRecorder()
    handlers.RegisterHandler(rr, req)
    
    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("expected %d, got %d", http.StatusCreated, status)
    }
}

func TestLoginSuccess(t *testing.T) {
    reqBody := `{"email":"john@example.com","password":"password123"}`
    req, _ := http.NewRequest("POST", "/login", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    
    rr := httptest.NewRecorder()
    handlers.LoginHandler(rr, req)
    
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected %d, got %d", http.StatusOK, status)
    }
}
