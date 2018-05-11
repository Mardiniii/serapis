package models

// JWT model for API key requests
type JWT struct {
	Token string `json:"token"`
}

// SecretKey constant for JWT
const SecretKey string = "serapis-api-secrete-key"
