package models

// User database schema 
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	DOB      string `json:"date_of_birth"`
}