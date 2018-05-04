package main

import (
	"encoding/json"
	"net/http"
)

// Generic routes handlers

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("All systems reporting at 100%")
}
