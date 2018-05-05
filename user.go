package main

import "time"

// User model
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
}

// Users collection
type Users []User
