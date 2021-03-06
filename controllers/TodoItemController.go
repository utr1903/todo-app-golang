package controllers

import (
	"net/http"

	"github.com/todo-app-golang/commons"
	"github.com/todo-app-golang/services/todoitemmodule"
)

// TodoItemController : Controller for Todo Item Model
type TodoItemController struct {
	Base *Controller
}

// GetTodoItems : Handler for getting all todo items
func (c *TodoItemController) GetTodoItems(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &todoitemmodule.TodoItemService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	items, err := s.GetTodoItems(dto)

	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	result := commons.Success(items, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// CreateTodoItem : Handler for creating a new item
func (c *TodoItemController) CreateTodoItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &todoitemmodule.TodoItemService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	itemID, err := s.CreateTodoItem(dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	result := commons.Success(itemID, nil)
	c.Base.CreateResponse(&w, http.StatusOK, result)
}

// UpdateTodoItem : Handler for updating an existing item
func (c *TodoItemController) UpdateTodoItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &todoitemmodule.TodoItemService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	err := s.UpdateTodoItem(dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(&w, http.StatusOK, nil)
}

// DeleteTodoItem : Handler for deleting an existing item
func (c *TodoItemController) DeleteTodoItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToString(&w, r)

	s := &todoitemmodule.TodoItemService{
		Req: r,
		Cu: &commons.CommonUtils{
			Db: c.Base.Db,
		},
		Db: c.Base.Db,
	}

	err := s.DeleteTodoItem(c.Base.Db, dto)
	if err != nil {
		c.Base.CreateResponse(&w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(&w, http.StatusOK, nil)
}
