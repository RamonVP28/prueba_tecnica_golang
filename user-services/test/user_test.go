package test

import (
	"prueba_tecnica_golang/user-services/utils"
	"testing"
)

// TestValidateEmail tests the email validation function with various inputs.
func TestValidateEmail(t *testing.T) {
	// Define test cases with expected results
	testCases := []struct {
		email    string
		expected bool
		msg      string
	}{
		{"example@gmail.com", true, "Standard email"},
		{"bademail.com", false, "Missing @ symbol"},
		{"bademail@com", false, "Missing domain suffix"},
		{"@gmail.com", false, "Missing local part"},
		{"name@.com", false, "Missing domain name"},
		{"name@gmail.com.", false, "Trailing dot"},
		{"", false, "Empty email"},
		{"name@localhost", false, "Localhost domain"},
		{"name@domain.c", false, "Single character domain suffix"},
		{"very.long.email.address@example.com", true, "Long email address"},
	}

	// Execute tests
	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			result := utils.ValidateEmail(tc.email)
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %v, got %v for email '%s'", tc.msg, tc.expected, result, tc.email)
			}
		})
	}
}
