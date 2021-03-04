package controllers

import (
	"log"
	"net/http"

	"github.com/todo-app-golang/services"
)

// UsersController : Controller for User Model
type UsersController struct {
	Base *Controller
}

// GetUsers : Handler for getting all users
func (c *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {

	// c.Base.ParseRequest(w, r)
	s := &services.UserService{}

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

	s := &services.UserService{}

	user, err := s.GetUser(c.Base.Db, userID)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, user)
}

// CreateUser : Handler for creating new user
func (c *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &services.UserService{}

	userID, err := s.CreateUser(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, userID)
}

// UpdateUser : Handler for updating an existing user
func (c *UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &services.UserService{}

	err := s.UpdateUser(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}

// DeleteUser : Handler for deleting an existing user
func (c *UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &services.UserService{}

	err := s.DeleteUser(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}
