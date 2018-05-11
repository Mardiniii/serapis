package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CreateEvaluation handler to process a new evaluation request
func CreateEvaluation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// if vars["platform"] != nil {
	//
	// }
	// RespondWithError(w, http.StatusNotFound, "Not implemented endpoint")

	RespondWithJSON(w, http.StatusOK, vars)
}
