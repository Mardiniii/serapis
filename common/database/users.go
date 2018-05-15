package database

import (
	"log"

	"github.com/Mardiniii/serapis/common/models"
	"github.com/pkg/errors"
)

func (conn *Postgres) createUsersTable() (err error) {
	if _, err = conn.Db.Exec(usersTable); err != nil {
		err = errors.Wrapf(err, "Can not create users table (%s)", usersTable)
		return
	}

	return
}

// CreateUser adds a new user record to the database
func (conn *Postgres) CreateUser(u *models.User) (err error) {
	err = conn.Db.QueryRow(createUser, u.Email, u.APIKey).Scan(&u.ID)
	if err != nil {
		log.Println(err)
	}
	return
}

// DeleteUser removes a new user record from the database
func (conn *Postgres) DeleteUser(id int) (err error) {
	_, err = conn.Db.Exec(deleteUser, 1)
	return
}

// FindUserByEmail queries for a user with the given email
func (conn *Postgres) FindUserByEmail(email string) (u models.User, err error) {
	row := conn.Db.QueryRow(userByEmail, email)

	err = row.Scan(&u.ID, &u.Email, &u.APIKey, &u.CreatedAt)
	return
}
