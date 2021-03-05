package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/todo-app-golang/commons"
	"github.com/todo-app-golang/services/usersmodule"
)

// UsersController : Controller for User Model
type UsersController struct {
	Base *Controller
}

// SignIn : Handler for signing in
func (c *UsersController) SignIn(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Get the JSON body and decode into credentials
	var creds commons.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&commons.Exception{M: commons.RequestNotValid})
		return
	}

	// Check given username and password
	s := &usersmodule.UserService{}
	userID, err := s.CheckUser(c.Base.Db, &creds.UserName, &creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&commons.Exception{M: commons.UserNotFound})
		return
	}

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &commons.Claims{
		UserID:   *userID,
		UserName: creds.UserName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(commons.JwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&commons.Exception{M: commons.TokenNotValid})
		return
	}

	// Set the client token
	result := &commons.Token{
		Token:      tokenString,
		ExpireDate: expirationTime,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetUsers : Handler for getting all users
func (c *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {

	// c.Base.ParseRequest(w, r)
	s := &usersmodule.UserService{}

	users, err := s.GetUsers(c.Base.Db)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, users)
}

// GetUser : Handler for getting user with given ID
func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToMap(w, r)
	userID, ok := dto["userId"].(string)

	w.Header().Set("Content-Type", "application/json")
	if !ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&commons.Exception{M: commons.UserIDNotValid})
	}

	s := &usersmodule.UserService{}

	user, err := s.GetUser(c.Base.Db, userID)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, user)
}

// CreateUser : Handler for creating new user
func (c *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &usersmodule.UserService{}

	userID, err := s.CreateUser(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, userID)
}

// UpdateUser : Handler for updating an existing user
func (c *UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &usersmodule.UserService{}

	err := s.UpdateUser(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}

// DeleteUser : Handler for deleting an existing user
func (c *UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &usersmodule.UserService{}

	err := s.DeleteUser(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}
