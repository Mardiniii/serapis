package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	dbs "github.com/Mardiniii/serapis_api/dbs"
	models "github.com/Mardiniii/serapis_api/models"
)

// GetAPIKey returns the JSON WEB Token for the user
func GetAPIKey(w http.ResponseWriter, r *http.Request) {
	// Extract query string param
	keys, params := r.URL.Query()["email"]
	email := string(keys[0])
	if !params || len(keys) < 1 || email == "" {
		err := errors.New("Email param is missing")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := dbs.RepoFindUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	t := models.JWT{Token: user.APIKey}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		log.Panicln(err)
	}
}
