package controllers

import (
	"log"
	"net/http"

	"github.com/todo-app-golang/services"
)

// TodoListController : Controller for Todo List Model
type TodoListController struct {
	Base *Controller
}

// GetTodoLists : Handler for all todo lists
func (c *TodoListController) GetTodoLists(w http.ResponseWriter, r *http.Request) {
	s := &services.TodoListService{}
	users, err := s.GetLists(c.Base.Db)

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, users)
}

// GetTodoList : Handler for getting a list with given ID
func (c *TodoListController) GetTodoList(w http.ResponseWriter, r *http.Request) {

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

// CreateTodoList : Handler for creating a new list
func (c *TodoListController) CreateTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &services.TodoListService{}

	listID, err := s.CreateTodoList(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, listID)
}

// UpdateTodoList : Handler for updating an existing list
func (c *TodoListController) UpdateTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &services.TodoListService{}

	err := s.UpdateTodoList(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}

// DeleteTodoList : Handler for deleting an existing list
func (c *TodoListController) DeleteTodoList(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &services.TodoListService{}

	err := s.DeleteTodoList(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}