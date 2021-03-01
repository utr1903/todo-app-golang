package controllers

import (
	"net/http"

	"github.com/todo-app-golang/services"
)

// TodoItemController : Controller for Todo Item Model
type TodoItemController struct {
	Base *Controller
}

// GetItems : Handler for todo items
func (c *TodoItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	s := &services.TodoItemService{}
	users, err := s.GetItems(c.Base.Db)

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, users)
}
