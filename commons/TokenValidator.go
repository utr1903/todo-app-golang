package commons

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// TokenValidator : Parser for tokens
type TokenValidator struct{}

// Claims : Claims for logged in user
type Claims struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

var jwtKey = []byte("some_dope_secret_key")

// ParseUserID : Parses UserId from token for further purposes
func ParseUserID(r *http.Request) (*string, error) {

	authorizationHeader := r.Header.Get("authorization")
	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) != 2 {
		return nil, errors.New("problem occured")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
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

		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {

			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {

				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})

				if err != nil {
					// json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}

				if token.Valid {
					// context.Set(req, "decoded", token.Claims)
					next(w, r)
				} else {
					// json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			// json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			return
		}
	})
}
