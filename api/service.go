package api

import (
	"log"
	"net/http"

	"github.com/Mardiniii/serapis/api/common"
	"github.com/Mardiniii/serapis/api/middlewares"
	"github.com/Mardiniii/serapis/api/routes"
	"github.com/Mardiniii/serapis/common/database"
	"github.com/urfave/negroni"
)

// Init starts API server
func Init() {
	pg := database.Connection()
	defer pg.Db.Close()
	var router = routes.Router()

	n := negroni.New()
	n.Use(negroni.HandlerFunc(middlewares.Logger))
	n.Use(negroni.HandlerFunc(middlewares.AuthHeaderValidator))
	n.Use(negroni.HandlerFunc(middlewares.Authenticator))
	n.UseHandler(router)

	println("Creating seed data")
	common.RepoSeedData()

	println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", n))
}
