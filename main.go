package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var router = NewRouter()

	fmt.Println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
