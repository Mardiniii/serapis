package main

import (
	"log"
	"net/http"

	dbs "./dbs"
	routes "./routes"
)

func main() {
	var router = routes.Router()

	println("Starting server on port: 8080")
	dbs.RepoSeedData()
	println("Seed data created")
	log.Fatal(http.ListenAndServe(":8080", router))
}
