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

// LoginUser handles user login.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Error processing data", http.StatusBadRequest)
		return
	}

	// Check for missing fields
	if credentials.Login == "" {
		http.Error(w, "Missing field: Username or Email", http.StatusBadRequest)
		return
	}

	if credentials.Password == "" {
		http.Error(w, "Missing field: Password", http.StatusBadRequest)
		return
	}

	// Find the user with the given login details
	var user models.User
	for _, usr := range users {
		if (usr.Email == credentials.Login || usr.Username == credentials.Login) && usr.Password == credentials.Password {
			user = usr
			break
		}
	}

	if user.Username == "" {
		http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
