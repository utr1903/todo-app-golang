package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/todo-app-golang/services/usersmodule"
)

// UsersController : Controller for User Model
type UsersController struct {
	Base *Controller
}

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

var jwtKey = []byte("some_dope_secret_key")

// SignIn : Handler for signing in
func (c *UsersController) SignIn(w http.ResponseWriter, r *http.Request) {

	// Get the JSON body and decode into credentials
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check given username and password
	s := &usersmodule.UserService{}

	if !s.CheckUserPassword(c.Base.Db, &creds.UserName) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserName: creds.UserName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	w.Write([]byte(cookie.String()))
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

	if !ok {
		log.Fatal("UserId is not valid")
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
