package models

import (
	"time"
)

// Evaluation model
type Evaluation struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	Platform  string    `json:"platform"`
	Request   string    `json:"request"`
	ExitCode  int       `json:"exit_code"`
	CreatedAt time.Time `json:"created_at"`
}
