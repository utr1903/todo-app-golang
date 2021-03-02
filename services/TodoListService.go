package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

// TodoList : TodoList model
type TodoList struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}

// TodoListService : Implementation of TodoListService
type TodoListService struct{}

// GetLists : Returns all lists
func (s *TodoListService) GetLists(db *sql.DB) ([]TodoList, error) {
	rows, err := db.Query("select * from lists")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lists := []TodoList{}

	for rows.Next() {
		var list TodoList
		if rows.Scan(&list.ID, &list.UserID, &list.Name) != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

// GetList : Returns list with given ID
func (s *TodoListService) GetList(db *sql.DB, itemID string) (*TodoList, error) {
	list := &TodoList{}

	q := "select * from lists where id = ?"
	err := db.QueryRow(q, itemID).
		Scan(&list.ID, &list.UserID, &list.Name)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// CreateTodoList : Creates a new todo list
func (s *TodoListService) CreateTodoList(db *sql.DB, dto *string) (*string, error) {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "insert into lists (id, userid, name) values (?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

	todoList := &TodoList{}
	json.Unmarshal([]byte(*dto), &todoList)

	listID := uuid.New().String()
	res, err := stmt.ExecContext(ctx, listID, todoList.UserID, todoList.Name)
	if err != nil {
		return nil, err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return nil, err
	}

	return &listID, nil
}
