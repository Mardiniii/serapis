package database

import (
	"database/sql"
	"fmt"
	"log"

	// Postgres package
	_ "github.com/lib/pq"
)

// Kubernetes connection
const (
	host     = "postgres-postgresql"
	port     = 5432
	user     = "admin"
	password = "hola-mama"
	dbname   = "serapis"
)

// Local connection
// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = ""
// 	dbname   = "serapis"
// )

var singletonConnection *Postgres

// Postgres creates a new postgres connection
type Postgres struct {
	Db *sql.DB
}

// Connection starts a connection to postgres
func Connection() *Postgres {
	if singletonConnection == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

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

// RunMigrations to create tables is don't exist yet
func (conn *Postgres) RunMigrations() {
	err := conn.createUsersTable()
	if err != nil {
		log.Print(err)
	}
	err = conn.createEvaluationsTable()
	if err != nil {
		log.Print(err)
	}
}
