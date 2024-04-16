package models

// Login represent the structure of a login
type Credentials struct {
	Login    string `json:"username"`
	Password string `json:"password"`
}
