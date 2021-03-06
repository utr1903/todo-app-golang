package commons

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenValidator : Parser for tokens
type TokenValidator struct{}

// Credentials : Username and password for signing in
type Credentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// Claims : Claims for logged in user
type Claims struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

// Token : Token for logged in user
type Token struct {
	Token      string    `json:"token"`
	ExpireDate time.Time `json:"expireDate"`
}

// JwtKey : Key for token generation
var JwtKey = []byte("some_dope_secret_key")

// ParseUserID : Parses UserId from token for further purposes
func ParseUserID(r *http.Request) (*string, error) {

	authorizationHeader := r.Header.Get("authorization")
	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) != 2 {
		return nil, errors.New("problem occured")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("problem occured")
	}

	return &claims.UserID, nil
}

// ValidateToken : Middleware for token authentication
func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {

			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {

				claims := &Claims{}
				token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
					return JwtKey, nil
				})

				if err != nil {
					// json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				if token.Valid {
					// context.Set(req, "decoded", token.Claims)
					next(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					// json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			// json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
