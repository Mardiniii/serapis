package dbs

import (
	"fmt"
	"time"

	models "../models"
)

// Users collection
type Users []models.User

var currentID int
var UsersRepo Users

// RepoSeedData gives some initial data
func RepoSeedData() {
	seeds := Users{
		models.User{Email: "email1@correo.com"},
		models.User{Email: "email2@correo.com"},
		models.User{Email: "email3@correo.com"},
		models.User{Email: "email4@correo.com"},
		models.User{Email: "email5@correo.com"},
	}

	for _, u := range seeds {
		RepoCreateUser(u)
	}
}

// RepoFindUser creates a new user
func RepoFindUser(id int) models.User {
	for _, u := range UsersRepo {
		if u.ID == id {
			return u
		}
	}

	// Not User found it
	return models.User{}
}

// RepoCreateUser creates a new user
func RepoCreateUser(u models.User) models.User {
	currentID++
	u.ID = currentID
	u.CreatedAt = time.Now()
	u.GenerateAPIKey()

	UsersRepo = append(UsersRepo, u)

	return u
}

// RepoDestroyUser creates a new user
func RepoDestroyUser(id int) error {
	for i, u := range UsersRepo {
		if u.ID == id {
			UsersRepo = append(UsersRepo[:i], UsersRepo[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find User with id %d to be deleted", id)
}
