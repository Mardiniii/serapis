package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	dbs "../dbs"
	models "../models"
)

// User routes handlers

// UserIndex handler for /users
func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(dbs.UsersRepo)
	if err != nil {
		log.Fatal(err)
	}
}

// UserCreate handler for user/ - POST
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Extract JSON payload
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	// Parse JSON data with User struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = json.NewEncoder(w).Encode("The user was not created")

		if err != nil {
			log.Fatal(err)
		}
	}

	// Create the new user
	u := dbs.RepoCreateUser(user)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		log.Fatal(err)
	}
}
