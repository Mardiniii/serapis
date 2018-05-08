package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	dbs "github.com/Mardiniii/serapis_api/dbs"
	models "github.com/Mardiniii/serapis_api/models"
)

// User routes handlers

// UserIndex handler for /users
func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(dbs.UsersRepo)
	if err != nil {
		log.Panicln(err)
	}
}

// UserCreate handler for user/ - POST
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error

	// Extract JSON payload
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Panicln(err)
	}
	defer r.Body.Close()

	// Parse JSON data with User struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = json.NewEncoder(w).Encode("The user was not created")

		if err != nil {
			log.Panicln(err)
		}
	}

	// Create the new user
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	user, err = dbs.RepoCreateUser(user)

	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Panicln(err)
		}
	}
}
