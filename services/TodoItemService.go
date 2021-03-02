package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

// TodoItem : TodoList model
type TodoItem struct {
	ID      string `json:"id"`
	ListID  string `json:"listId"`
	Content string `json:"content"`
}

// TodoItemService : Implementation of TodoItemService
type TodoItemService struct{}

// GetItems : Returns all items
func (s *TodoItemService) GetItems(db *sql.DB) ([]TodoItem, error) {
	rows, err := db.Query("select * from items")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []TodoItem{}

	for rows.Next() {
		var item TodoItem
		if rows.Scan(&item.ID, &item.ListID, &item.Content) != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// GetItem : Returns item with given ID
func (s *TodoItemService) GetItem(db *sql.DB, itemID string) (*TodoItem, error) {
	item := &TodoItem{}

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

	todoItem := &TodoItem{}
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

func (s *TodoItemService) doesListExist(db *sql.DB, listID *string) bool {
	q := "select id from lists where id = ?"
	err := db.QueryRow(q, listID).Scan()
	if err != nil {
		return false
	}

	return true
}
