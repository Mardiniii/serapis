package api

import (
	"log"
	"net/http"

	"github.com/Mardiniii/serapis/api/middlewares"
	"github.com/Mardiniii/serapis/api/routes"
	db "github.com/Mardiniii/serapis/common/database"
	"github.com/urfave/negroni"
)

// Init starts API server
func Init() {
	pg := db.Connection()
	pg.RunMigrations()

	defer pg.Db.Close()
	var router = routes.Router()

	n := negroni.New()
	n.Use(negroni.HandlerFunc(middlewares.Logger))
	n.Use(negroni.HandlerFunc(middlewares.AuthHeaderValidator))
	n.Use(negroni.HandlerFunc(middlewares.Authenticator))
	n.UseHandler(router)

	println("Creating seed data")
	db.RepoSeedData()

	println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", n))
}
