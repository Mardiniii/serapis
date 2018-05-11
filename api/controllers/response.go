package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithError function
func RespondWithError(w http.ResponseWriter, code int, err string) {
	RespondWithJSON(w, code, map[string]string{"error": err})
}

// RespondWithJSON function
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Panicln(err)
	}
}
