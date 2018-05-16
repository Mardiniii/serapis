package models

import (
	"time"
)

// Evaluation model
type Evaluation struct {
	ID           int         `json:"id"`
	UserID       int         `json:"user_id"`
	Status       string      `json:"status"`
	Language     string      `json:"language"`
	Code         string      `json:"code"`
	Stdin        []string    `json:"stdin"`
	Dependencies PropertyMap `json:"dependencies"`
	Git          PropertyMap `json:"git"`
	Output       string      `json:"output"`
	ExitCode     int         `json:"exit_code"`
	CreatedAt    time.Time   `json:"created_at"`
}
