package controllers

import (
	"net/http"

	"github.com/todo-app-golang/services"
)

// UsersController : Controller for User Model
type UsersController struct {
	Base *Controller
}

// GetUsers : Handler for users
func (c *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	s := &services.UserService{}
	users, err := s.GetUsers(c.Base.Db)

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, users)
}
