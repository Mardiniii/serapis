package models

import (
	"encoding/json"
	"time"
)

// Evaluation model
type Evaluation struct {
	ID           int             `json:"id"`
	UserID       int             `json:"user_id"`
	Status       string          `json:"status"`
	Language     string          `json:"language"`
	Code         string          `json:"code"`
	Stdin        []string        `json:"stdin"`
	Dependencies json.RawMessage `json:"dependencies"`
	Git          json.RawMessage `json:"git"`
	Output       string          `json:"output"`
	ExitCode     int             `json:"exit_code"`
	CreatedAt    time.Time       `json:"created_at"`
}
