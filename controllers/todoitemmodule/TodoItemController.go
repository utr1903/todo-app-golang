package todoitemmodule

import (
	"log"
	"net/http"

	"github.com/todo-app-golang/controllers"
	"github.com/todo-app-golang/services"
)

// TodoItemController : Controller for Todo Item Model
type TodoItemController struct {
	Base *controllers.Controller
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

// GetItem : Handler for getting item with given ID
func (c *TodoItemController) GetItem(w http.ResponseWriter, r *http.Request) {

	dto := c.Base.ParseRequestToMap(w, r)
	itemID, ok := dto["itemId"].(string)

	if !ok {
		log.Fatal("ItemID is not valid")
	}

	s := &services.TodoItemService{}

	item, err := s.GetItem(c.Base.Db, itemID)
	if err != nil {
		c.Base.CreateResponse(w, http.StatusBadRequest, nil)
	}

	c.Base.CreateResponse(w, http.StatusOK, item)
}
