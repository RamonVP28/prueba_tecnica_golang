package handlers

import (
	"encoding/json"
	"net/http"
	"prueba_tecnica_golang/user-services/models"
	"prueba_tecnica_golang/user-services/utils"
)

// in-memory store to simulate a database
var users = make(map[string]models.User)

// CheckUserExists checks if the user is already registered
func CheckUserExists(email, phone, username string) bool {
	for _, user := range users {
		if user.Email == email || user.Phone == phone || user.Username == username {
			return true
		}
	}
	return false
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error processing data", http.StatusBadRequest)
		return
	}

	// Check for missing fields
	if user.Username == "" {
		http.Error(w, "Missing field: Username", http.StatusBadRequest)
		return
	}
	if user.Email == "" {
		http.Error(w, "Missing field: Email", http.StatusBadRequest)
		return
	}
	if user.Phone == "" {
		http.Error(w, "Missing field: Phone", http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		http.Error(w, "Missing field: Password", http.StatusBadRequest)
		return
	}

	if !utils.ValidateEmail(user.Email) || !utils.ValidatePhone(user.Phone) || !utils.ValidatePassword(user.Password) {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	if CheckUserExists(user.Email, user.Phone, user.Username) {
		http.Error(w, "Email, Username or phone already registered", http.StatusBadRequest)
		return
	}
	// Simulate user storage
	users[user.Email] = user
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered successfully"))
}
