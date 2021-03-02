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

	todoList := &TodoList{}
	json.Unmarshal([]byte(*dto), &todoList)

	// Check whether to be created list is assigning to an existing user
	if !s.doesUserExist(db, &todoList.UserID) {
		return nil, nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "insert into lists (id, userid, name) values (?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

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

// UpdateTodoList : Updates an existing todo list
func (s *TodoListService) UpdateTodoList(db *sql.DB, dto *string) error {

	todoList := &TodoList{}
	json.Unmarshal([]byte(*dto), &todoList)

	// Check whether to be created list is assigning to an existing user
	if !s.doesListExist(db, &todoList.ID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update lists set name = ? where id = ?"
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, todoList.Name, todoList.ID)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}

func (s *TodoListService) doesUserExist(db *sql.DB, userID *string) bool {
	q := "select id from users where id = ?"
	var userExists string
	err := db.QueryRow(q, userID).Scan(&userExists)
	if err != nil {
		return false
	}

	return true
}

func (s *TodoListService) doesListExist(db *sql.DB, listID *string) bool {
	q := "select id from lists where id = ?"
	var listExists string
	err := db.QueryRow(q, listID).Scan(&listExists)
	if err != nil {
		return false
	}

	return true
}
