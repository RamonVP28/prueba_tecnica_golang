package test

import (
	"prueba_tecnica_golang/user-services/utils"
	"testing"
)

func TestValidatePhone(t *testing.T) {
	testCases := []struct {
		phone    string
		expected bool
		msg      string
	}{
		{"1234567890", true, "Valid phone number"},
		{"123456789", false, "Too short"},
		{"12345678901", false, "Too long"},
		{"abcdefghij", false, "Non-numeric characters"},
		{"123 456 7890", false, "Spaces included"},
		{"(123)4567890", false, "Parentheses included"},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			result := utils.ValidatePhone(tc.phone)
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %v, got %v for phone '%s'", tc.msg, tc.expected, result, tc.phone)
			}
		})
	}
}
