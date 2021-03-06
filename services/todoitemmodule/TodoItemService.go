package todoitemmodule

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/todo-app-golang/commons"
	"github.com/todo-app-golang/services/todoitemmodule/dtos"
)

// TodoItemService : Implementation of TodoItemService
type TodoItemService struct {
	Req *http.Request
	Db  *sql.DB
	Cu  *commons.CommonUtils
}

// GetTodoItems : Returns all items in the given list of user
func (s *TodoItemService) GetTodoItems(listID *string) ([]dtos.GetTodoItems, error) {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return nil, err
	}

	// Check whether caller user exists
	if !s.Cu.DoesUserExist(s.Db, userID) {
		return nil, nil
	}

	q := "select Id, Content from items where ListId = ?"

	rows, err := s.Db.Query(q, listID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []dtos.GetTodoItems{}

	for rows.Next() {
		var item dtos.GetTodoItems
		if rows.Scan(&item.ItemID, &item.Content) != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// CreateTodoItem : Creates a new todo item
func (s *TodoItemService) CreateTodoItem(dto *string) (*string, error) {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return nil, err
	}

	// Check whether caller user exists
	if !s.Cu.DoesUserExist(s.Db, userID) {
		return nil, nil
	}

	todoItem := &dtos.TodoItem{}
	json.Unmarshal([]byte(*dto), &todoItem)

	// Check whether to be created list is assigning to an existing user
	if !s.Cu.DoesListExist(s.Db, &todoItem.ListID, userID) {
		return nil, nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "insert into items (id, listid, content) values (?, ?, ?)"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

	itemID := uuid.New().String()
	res, err := stmt.ExecContext(ctx, itemID, todoItem.ListID, todoItem.Content)
	if err != nil {
		return nil, err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return nil, err
	}

	return &itemID, nil
}

// UpdateTodoItem : Updates an existing todo item
func (s *TodoItemService) UpdateTodoItem(dto *string) error {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return err
	}

	// Check whether caller user exists
	if !s.Cu.DoesUserExist(s.Db, userID) {
		return nil
	}

	todoItem := &dtos.TodoItem{}
	json.Unmarshal([]byte(*dto), &todoItem)

	// Check whether to be updated item exists and belongs to the caller
	if !s.Cu.DoesItemExist(s.Db, &todoItem.ID, userID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update items set content = ? where id = ?"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, todoItem.Content, todoItem.ID)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}

// DeleteTodoItem : Deletes an existing item
func (s *TodoItemService) DeleteTodoItem(db *sql.DB, itemID *string) error {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return err
	}

	// Check whether caller user exists
	if !s.Cu.DoesUserExist(db, userID) {
		return nil
	}

	// Check whether to be updated item exists and belongs to the caller
	if !s.Cu.DoesItemExist(db, itemID, userID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "delete from items where id = ?"
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, itemID)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}
