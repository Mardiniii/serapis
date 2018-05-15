package database

import (
	"database/sql"
	"fmt"
	"log"

	// Postgres package
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "serapis"
)

var singletonConnection *Postgres

// Postgres creates a new postgres connection
type Postgres struct {
	Db *sql.DB
}

// Connection starts a connection to postgres
func Connection() *Postgres {
	if singletonConnection == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
			host, port, user, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		// sql.Open only checks params are valid with
		// the Ping we will create the connection
		err = db.Ping()
		if err != nil {
			panic(err)
		}

		singletonConnection = &Postgres{db}
		log.Println("Successfully connected!")
	}

	return singletonConnection
}

func (conn *Postgres) createEvaluationsTable() (err error) {
	if _, err = conn.Db.Exec(evaluationsTable); err != nil {
		err = errors.Wrapf(err, "Can not create evaluations table (%s)", evaluationsTable)
		return
	}

	return
}
