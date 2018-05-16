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
	err = conn.Db.QueryRow(createUser, u.Email, u.APIKey).Scan(&u.ID, &u.CreatedAt)
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

// GetUsers returns all the users stored in database
func (conn *Postgres) GetUsers() (users []models.User, err error) {
	rows, err := conn.Db.Query(allUsers)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u = models.User{}
		err = rows.Scan(&u.ID, &u.Email, &u.APIKey, &u.CreatedAt)
		users = append(users, u)
		if err != nil {
			log.Println(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return
}
