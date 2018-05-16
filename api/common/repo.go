package common

import (
	db "github.com/Mardiniii/serapis/common/database"
	"github.com/Mardiniii/serapis/common/models"
)

// RepoSeedData gives some initial data
func RepoSeedData() {
	seeds := []models.User{
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

	u.APIKey = GenerateAPIKey(u)
	err = db.Connection().CreateUser(&u)
	return u, err
}

// RepoDestroyUser removes a user
func RepoDestroyUser(id int) (deleted bool, err error) {
	err = db.Connection().DeleteUser(id)
	return
}

// RepoUsers returns all the users
func RepoUsers() (users []models.User) {
	users, _ = db.Connection().GetUsers()
	return
}

// RepoCreateEvaluation creates a new evaluation
func RepoCreateEvaluation(eval *models.Evaluation) error {
	var err error

	err = db.Connection().CreateEvaluation(eval)
	return err
}
