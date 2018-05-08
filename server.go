package main

import (
	"log"
	"net/http"

	dbs "github.com/Mardiniii/serapis_api/dbs"
	middlewares "github.com/Mardiniii/serapis_api/middlewares"
	routes "github.com/Mardiniii/serapis_api/routes"
	"github.com/urfave/negroni"
)

func main() {
	var router = routes.Router()

	n := negroni.New()
	n.Use(negroni.HandlerFunc(middlewares.Logger))
	n.UseHandler(router)

	println("Creating seed data")
	dbs.RepoSeedData()

	println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", n))
}
