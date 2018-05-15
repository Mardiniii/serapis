package common

import (
	"time"

	db "github.com/Mardiniii/serapis/common/database"
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
func RepoFindUserByEmail(email string) (u models.User, err error) {
	u, err = db.Connection().FindUserByEmail(email)
	return
}

// RepoCreateUser creates a new user
func RepoCreateUser(u models.User) (models.User, error) {
	var err error

	currentID++
	u.ID = currentID
	u.CreatedAt = time.Now()
	u.APIKey = GenerateAPIKey(u)
	err = db.Connection().CreateUser(&u)
	return u, err
}

// RepoDestroyUser removes a user
func RepoDestroyUser(id int) (deleted bool, err error) {
	err = db.Connection().DeleteUser(id)
	return
}
