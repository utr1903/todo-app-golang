package todolistmodule

import (
	"log"
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

// GetList : Handler for getting list with given ID
func (c *TodoListController) GetList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToMap(w, r)
	listID, ok := dto["listId"].(string)

	if !ok {
		log.Fatal("ListID is not valid")
	}

	s := &services.TodoListService{}

	list, err := s.GetList(c.Base.Db, listID)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, list)
}
