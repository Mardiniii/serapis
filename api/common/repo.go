package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/Mardiniii/serapis/common/models"
)

// Users collection
type Users []models.User

var currentID int

// UsersRepo is a temporary users storage
var UsersRepo Users

func uniqueEmail(email string) bool {
	for _, u := range UsersRepo {
		if u.Email == email {
			return false
		}
	}

	return true
}

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

// RepoFindUserByEmail finds a new user matching the given email
func RepoFindUserByEmail(email string) (models.User, error) {
	for _, u := range UsersRepo {
		if u.Email == email {
			return u, nil
		}
	}
	err := errors.New("User not found")
	// Not User found it
	return models.User{}, err
}

// RepoCreateUser creates a new user
func RepoCreateUser(u models.User) (models.User, error) {
	var err error

	if uniqueEmail(u.Email) {
		currentID++
		u.ID = currentID
		u.CreatedAt = time.Now()
		u.APIKey = GenerateAPIKey(u)
		UsersRepo = append(UsersRepo, u)
		return u, nil
	}
	err = errors.New("Email already exists")
	return models.User{}, err
}

// RepoDestroyUser creates a new user
func RepoDestroyUser(id int) (bool, error) {
	for i, u := range UsersRepo {
		if u.ID == id {
			UsersRepo = append(UsersRepo[:i], UsersRepo[i+1:]...)
			return true, nil
		}
	}
	err := fmt.Errorf("Could not find User with id %d to be deleted", id)

	return false, err
}
