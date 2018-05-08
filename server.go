package main

import (
	"log"
	"net/http"

	dbs "github.com/Mardiniii/serapis_api/dbs"
	routes "github.com/Mardiniii/serapis_api/routes"
)

func main() {
	var router = routes.Router()

	println("Starting server on port: 8080")
	dbs.RepoSeedData()
	println("Seed data created")
	log.Fatal(http.ListenAndServe(":8080", router))
}
