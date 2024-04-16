package handlers

import (
	"encoding/json"
	"net/http"
	"prueba_tecnica_golang/user-services/models"
	"prueba_tecnica_golang/user-services/utils"
)

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
