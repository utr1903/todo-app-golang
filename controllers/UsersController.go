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
		json.NewEncoder(w).Encode(commons.RequestNotValid())
		return
	}

	// Check given username and password
	s := &usersmodule.UserService{
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	userID, err := s.CheckUser(&creds.UserName, &creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(commons.UserNotFound())
		return
	}

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(1 * time.Hour)

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
		json.NewEncoder(w).Encode(commons.TokenNotValid())
		return
	}

	// Set the client token
	data := &commons.Token{
		Token:      tokenString,
		ExpireDate: expirationTime,
	}

	result := commons.Success(data, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// GetUser : Handler for getting user with given ID
func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToMap(&w, r)
	userID, ok := dto["userId"].(string)

	if !ok {
		result := commons.UserIDNotValid()
		c.Base.CreateResponse(&w, http.StatusOK, result)
	}

	s := &usersmodule.UserService{
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	user, err := s.GetUser(userID)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	result := commons.Success(user, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// CreateUser : Handler for creating new user
func (c *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &usersmodule.UserService{
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	userID, err := s.CreateUser(dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusInternalServerError, nil)
	}

	result := commons.Success(userID, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// UpdateUser : Handler for updating an existing user
func (c *UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &usersmodule.UserService{
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	err := s.UpdateUser(dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(&w, http.StatusOK, nil)
}

// DeleteUser : Handler for deleting an existing user
func (c *UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &usersmodule.UserService{
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	err := s.DeleteUser(dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(&w, http.StatusOK, nil)
}
