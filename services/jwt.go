package services

import (
	"log"

	models "github.com/Mardiniii/serapis_api/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateAPIKey using JWT
func GenerateAPIKey(u models.User) string {
	var err error
	var apiKey string

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Email,
	})
	apiKey, err = token.SignedString([]byte(models.SecretKey))
	if err != nil {
		log.Panicln(err)
	}

	return apiKey
}
