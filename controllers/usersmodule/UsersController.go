package usersmodule

import (
	"net/http"

	"github.com/todo-app-golang/controllers"
	"github.com/todo-app-golang/services"
)

// UsersController : Controller for User Model
type UsersController struct {
	Base *controllers.Controller
}

// GetUsers : Handler for users
func (c *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {

	// c.Base.ParseRequest(w, r)
	s := &services.UserService{}
	users, err := s.GetUsers(c.Base.Db)

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, users)
}
