package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Logger logs out all the requests
func Logger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)

	next(w, r)
}
