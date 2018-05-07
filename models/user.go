package models

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// User model
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
}

// GenerateAPIKey using JSON WEB Token
func (u *User) GenerateAPIKey() {
	var err error

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Email,
	})

	u.APIKey, err = token.SignedString([]byte("serapis"))
	if err != nil {
		log.Fatal(err)
	}
}
