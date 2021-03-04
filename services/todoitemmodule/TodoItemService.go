package todoitemmodule

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/todo-app-golang/services/todoitemmodule/dtos"
)

// TodoItemService : Implementation of TodoItemService
type TodoItemService struct{}

// GetItems : Returns all items
func (s *TodoItemService) GetItems(db *sql.DB) ([]dtos.TodoItem, error) {
	rows, err := db.Query("select * from items")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []dtos.TodoItem{}

	for rows.Next() {
		var item dtos.TodoItem
		if rows.Scan(&item.ID, &item.ListID, &item.Content) != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// GetItem : Returns item with given ID
func (s *TodoItemService) GetItem(db *sql.DB, itemID string) (*dtos.TodoItem, error) {
	item := &dtos.TodoItem{}

	q := "select * from items where id = ?"
	err := db.QueryRow(q, itemID).
		Scan(&item.ID, &item.ListID, &item.Content)

	if err != nil {
		return nil, err
	}

	return item, nil
}

// CreateTodoItem : Creates a new todo item
func (s *TodoItemService) CreateTodoItem(db *sql.DB, dto *string) (*string, error) {

	todoItem := &dtos.TodoItem{}
	json.Unmarshal([]byte(*dto), &todoItem)

	// Check whether to be created list is assigning to an existing user
	if !s.doesListExist(db, &todoItem.ListID) {
		return nil, nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "insert into items (id, listid, content) values (?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, q)
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
func (s *TodoItemService) UpdateTodoItem(db *sql.DB, dto *string) error {

	todoItem := &dtos.TodoItem{}
	json.Unmarshal([]byte(*dto), &todoItem)

	// Check whether to be created list is assigning to an existing user
	if !s.doesItemExist(db, &todoItem.ID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update items set content = ? where id = ?"
	stmt, err := db.PrepareContext(ctx, q)
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
func (s *TodoItemService) DeleteTodoItem(db *sql.DB, dto *string) error {

	var itemID string
	json.Unmarshal([]byte(*dto), &itemID)

	// Check whether to be created list is assigning to an existing user
	if !s.doesItemExist(db, &itemID) {
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

func (s *TodoItemService) doesListExist(db *sql.DB, listID *string) bool {
	q := "select id from lists where id = ?"
	var listExists string
	err := db.QueryRow(q, listID).Scan(&listExists)
	if err != nil {
		return false
	}

	return true
}

func (s *TodoItemService) doesItemExist(db *sql.DB, itemID *string) bool {
	q := "select id from items where id = ?"
	var itemExists string
	err := db.QueryRow(q, itemID).Scan(&itemExists)
	if err != nil {
		return false
	}

	return true
}
