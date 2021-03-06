package controllers

import (
	"net/http"

	"github.com/todo-app-golang/commons"
	"github.com/todo-app-golang/services/todolistmodule"
)

// TodoListController : Controller for Todo List Model
type TodoListController struct {
	Base *Controller
}

// CreateTodoList : Handler for creating a new list
func (c *TodoListController) CreateTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &todolistmodule.TodoListService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}
	listID, err := s.CreateTodoList(dto)

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, listID)
}

// GetTodoLists : Handler for all todo lists
func (c *TodoListController) GetTodoLists(w http.ResponseWriter, r *http.Request) {

	s := &todolistmodule.TodoListService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	lists, err := s.GetTodoLists()

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, lists)
}

// UpdateTodoList : Handler for updating an existing list
func (c *TodoListController) UpdateTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &todolistmodule.TodoListService{Req: r}
	err := s.UpdateTodoList(dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}

// DeleteTodoList : Handler for deleting an existing list
func (c *TodoListController) DeleteTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &todolistmodule.TodoListService{Req: r}
	err := s.DeleteTodoList(dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}
