package dtos

import (
	"github.com/dgrijalva/jwt-go"
)

// User : User DB model
type User struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// Claims : Claims for token management
type Claims struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

// JwtKey : Secret key for JWT token
var JwtKey = []byte("my_secret_key")
