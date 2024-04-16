package utils

import (
	"regexp"
)

// ValidateEmail checks if the email is valid
func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

// ValidatePhone checks if the phone number is valid and has 10 digits
func ValidatePhone(phone string) bool {
	// Ensure the phone number is exactly 10 digits without spaces or any other characters
	return regexp.MustCompile(`^\d{10}$`).MatchString(phone)
}

// ValidatePassword checks if the password is valid and if is complex enough
func ValidatePassword(password string) bool {
	if len(password) < 6 || len(password) > 12 {
		return false
	}
	var (
		hasUpper   = false
		hasLower   = false
		hasDigit   = false
		hasSpecial = regexp.MustCompile(`[@$&]`).MatchString(password)
	)
	for _, c := range password {
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasDigit = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}
