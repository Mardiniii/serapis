package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	dbs "github.com/Mardiniii/serapis_api/dbs"
	models "github.com/Mardiniii/serapis_api/models"
)

// User routes handlers

// GetUsers handler for /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, dbs.UsersRepo)
}

// CreateUser handler for user/ - POST
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Extract JSON payload
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	// Parse JSON data with User struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, "The user was not created")
		return
	}

	// Create the new user
	user, err = dbs.RepoCreateUser(user)
	if err != nil {
		respondWithError(w, http.StatusConflict, "The user was not created")
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

// GetAPIKey returns the JSON WEB Token for the user
func GetAPIKey(w http.ResponseWriter, r *http.Request) {
	// Extract query string param
	keys, params := r.URL.Query()["email"]
	if !params || len(keys) < 1 || string(keys[0]) == "" {
		respondWithError(w, http.StatusBadRequest, "Email param not given in the query string")
		return
	}

	user, err := dbs.RepoFindUserByEmail(string(keys[0]))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	t := models.JWT{Token: user.APIKey}
	respondWithJSON(w, http.StatusCreated, t)
}
