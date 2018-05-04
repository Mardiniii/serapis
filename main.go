package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")

	fmt.Println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("All systems reporting at 100%")
}
