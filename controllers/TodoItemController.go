package controllers

import (
	"log"
	"net/http"

	"github.com/todo-app-golang/services/todoitemmodule"
)

// TodoItemController : Controller for Todo Item Model
type TodoItemController struct {
	Base *Controller
}

// GetTodoItems : Handler for getting all todo items
func (c *TodoItemController) GetTodoItems(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &todoitemmodule.TodoItemService{Req: r}
	items, err := s.GetTodoItems(c.Base.Db, dto)

	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, items)
}

// GetTodoItem : Handler for getting an item with given ID
func (c *TodoItemController) GetTodoItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToMap(w, r)
	itemID, ok := dto["itemId"].(string)

	if !ok {
		log.Fatal("ItemID is not valid")
	}

	s := &todoitemmodule.TodoItemService{Req: r}

	item, err := s.GetItem(c.Base.Db, itemID)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, item)
}

// CreateTodoItem : Handler for creating a new item
func (c *TodoItemController) CreateTodoItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &todoitemmodule.TodoItemService{Req: r}

	itemID, err := s.CreateTodoItem(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, itemID)
}

// UpdateTodoItem : Handler for updating an existing item
func (c *TodoItemController) UpdateTodoItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &todoitemmodule.TodoItemService{Req: r}

	err := s.UpdateTodoItem(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}

// DeleteTodoItem : Handler for deleting an existing item
func (c *TodoItemController) DeleteTodoItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(w, r)

	s := &todoitemmodule.TodoItemService{Req: r}

	err := s.DeleteTodoItem(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, nil)
}
