package todolistmodule

import (
	"net/http"

	"github.com/todo-app-golang/controllers"
	"github.com/todo-app-golang/services"
)

// TodoListController : Controller for Todo List Model
type TodoListController struct {
	Base *controllers.Controller
}

// GetLists : Handler for todo lists
func (c *TodoListController) GetLists(w http.ResponseWriter, r *http.Request) {
	s := &services.TodoListService{}
	users, err := s.GetLists(c.Base.Db)

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, users)
}
