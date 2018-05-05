package main

import (
	"fmt"
	"time"
)

var currentID int
var users Users

// RepoSeedData gives some initial data
func RepoSeedData() {
	seeds := Users{
		User{Email: "correo1@correo.com", APIKey: "API KEY 1"},
		User{Email: "correo2@correo.com", APIKey: "API KEY 2"},
		User{Email: "correo3@correo.com", APIKey: "API KEY 3"},
		User{Email: "correo4@correo.com", APIKey: "API KEY 4"},
		User{Email: "correo5@correo.com", APIKey: "API KEY 5"},
	}

	for _, u := range seeds {
		RepoCreateUser(u)
	}
}

// RepoFindUser creates a new user
func RepoFindUser(id int) User {
	for _, u := range users {
		if u.ID == id {
			return u
		}
	}

	// Not User found it
	return User{}
}

// RepoCreateUser creates a new user
func RepoCreateUser(u User) User {
	currentID++
	u.ID = currentID
	u.CreatedAt = time.Now()
	users = append(users, u)

	return u
}

// RepoDestroyUser creates a new user
func RepoDestroyUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find User with id %d to be deleted", id)
}
