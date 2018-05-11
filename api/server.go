package api

import (
	"log"
	"net/http"

	dbs "github.com/Mardiniii/serapis/api/dbs"
	middlewares "github.com/Mardiniii/serapis/api/middlewares"
	routes "github.com/Mardiniii/serapis/api/routes"
	"github.com/urfave/negroni"
)

// Init inits API server
func Init() {
	var router = routes.Router()

	n := negroni.New()
	n.Use(negroni.HandlerFunc(middlewares.Logger))
	n.Use(negroni.HandlerFunc(middlewares.AuthHeaderValidator))
	n.Use(negroni.HandlerFunc(middlewares.Authenticator))
	n.UseHandler(router)

	println("Creating seed data")
	dbs.RepoSeedData()

	println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", n))
}
