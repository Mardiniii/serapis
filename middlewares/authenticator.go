package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Mardiniii/serapis_api/controllers"
	"github.com/Mardiniii/serapis_api/services"
)

func extractHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

func extractAuthorizationHeader(r *http.Request) string {
	authHeader := extractHeader(r, "Authorization")

	return strings.Split(authHeader, " ")[1]
}

// AuthHeaderValidator middleware to check for API Key(JWT)
func AuthHeaderValidator(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var authHeader string
	var err error

	authHeader = extractHeader(r, "Authorization")
	values := strings.Split(authHeader, " ")
	if authHeader == "" {
		err = errors.New("An Authorization header wasn't given")
	} else if len(values) < 2 {
		err = errors.New("A wrong Authorization header was given")
	}
	if err != nil {
		controllers.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	next(w, r)
}

// Authenticator middleware for JWT
func Authenticator(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	apiKey := extractAuthorizationHeader(r)

	valid := services.ValidateAPIKey(apiKey)
	if !valid {
		controllers.RespondWithError(w, http.StatusUnauthorized, "Unvalid API Key")
		return
	}
	next(w, r)
}
