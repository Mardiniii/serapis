package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Generic routes handlers

// HealthCheck handler for /healthcheck
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("All systems reporting at 100%")
	if err != nil {
		log.Fatal(err)
	}
}
