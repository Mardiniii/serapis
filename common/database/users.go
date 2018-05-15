package database

import (
	"fmt"
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
	fmt.Println("Creando:", u.Email)
	err = conn.Db.QueryRow(createUser, u.Email, u.APIKey).Scan(&u.ID)
	if err != nil {
		log.Println(err)
	}

	return
}
