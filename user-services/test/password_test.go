package test

import (
	"prueba_tecnica_golang/user-services/utils"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
		msg      string
	}{
		{"Passw0rd@", true, "Valid password"},
		{"passw0rd@", false, "No uppercase letter"},
		{"PASSW0RD@", false, "No lowercase letter"},
		{"Password", false, "No digit"},
		{"Password1", false, "No special character"},
		{"P@ss", false, "Too short"},
		{"P@ssw0rddddd1", false, "Too long"},
		{"Passw0rd@", true, "Contains all required"},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			result := utils.ValidatePassword(tc.password)
			if result != tc.expected {
				t.Errorf("Failed %s: Expected %v, got %v for password '%s'", tc.msg, tc.expected, result, tc.password)
			}
		})
	}
}
