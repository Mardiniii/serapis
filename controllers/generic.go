package controllers

import (
	"net/http"
)

// Generic routes handlers

// HealthCheck handler for /healthcheck
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(
		w,
		http.StatusOK,
		map[string]string{"status": "All systems reporting at 100%"},
	)
}
