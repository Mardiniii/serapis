package services

import (
	"fmt"
	"log"

	models "github.com/Mardiniii/serapis/api/models"
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

// ValidateAPIKey using JWT
func ValidateAPIKey(apiKey string) bool {
	token, err := jwt.Parse(apiKey, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(models.SecretKey), nil
	})
	if err != nil {
		log.Println("Error", err.Error())
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], "user authenticated")
		return true
	}

	return false
}
