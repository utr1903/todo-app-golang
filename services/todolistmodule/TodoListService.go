package todolistmodule

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/todo-app-golang/commons"
	"github.com/todo-app-golang/services/todolistmodule/dtos"
)

// TodoListService : Implementation of TodoListService
type TodoListService struct {
	Req *http.Request
}

// CreateTodoList : Creates a new todo list
func (s *TodoListService) CreateTodoList(db *sql.DB, dto *string) (*string, error) {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return nil, err
	}

	// Check whether caller user exists
	if !s.doesUserExist(db, userID) {
		return nil, nil
	}

	todoList := &dtos.TodoList{
		UserID: *userID,
		Name:   *dto,
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "insert into lists (Id, UserId, Name) values (?, ?, ?)"
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

// GetTodoLists : Returns all lists
func (s *TodoListService) GetTodoLists(db *sql.DB) ([]dtos.GetTodoLists, error) {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return nil, err
	}

	// Check whether caller user exists
	if !s.doesUserExist(db, userID) {
		return nil, nil
	}

	q := "select Id, Name from lists where UserId = ?"

	rows, err := db.Query(q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lists := []dtos.GetTodoLists{}

	for rows.Next() {
		var list dtos.GetTodoLists
		if rows.Scan(&list.ListID, &list.ListName) != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

// GetTodoList : Returns list with given ID and it's items
func (s *TodoListService) GetTodoList(db *sql.DB, listID string) (*dtos.TodoList, error) {
	list := &dtos.TodoList{}

	q := "select * from items where ListId = ?"
	err := db.QueryRow(q, listID).
		Scan(&list.ID, &list.UserID, &list.Name)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// UpdateTodoList : Updates an existing todo list
func (s *TodoListService) UpdateTodoList(db *sql.DB, dto *string) error {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return err
	}

	// Check whether caller user exists
	if !s.doesUserExist(db, userID) {
		return nil
	}

	todoList := &dtos.TodoList{}
	json.Unmarshal([]byte(*dto), &todoList)

	// Check whether to be updated list exists and belongs to the caller
	if !s.doesListExist(db, &todoList.ID, userID) {
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

// DeleteTodoList : Deletes an existing list
func (s *TodoListService) DeleteTodoList(db *sql.DB, dto *string) error {

	userID, err := commons.ParseUserID(s.Req)
	if err != nil {
		return err
	}

	// Check whether caller user exists
	if !s.doesUserExist(db, userID) {
		return nil
	}

	var listID string
	json.Unmarshal([]byte(*dto), &listID)

	// Check whether to be updated list exists and belongs to the caller
	if !s.doesListExist(db, &listID, userID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "delete from lists where id = ?"
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, listID)
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
	q := "select Id from users where Id = ?"
	var userExists string
	err := db.QueryRow(q, userID).Scan(&userExists)
	if err != nil {
		return false
	}

	return true
}

func (s *TodoListService) doesListExist(db *sql.DB, listID *string, userID *string) bool {
	q := "select Id from lists where Id = ? and UserId = ?"
	var listExists string
	err := db.QueryRow(q, listID, userID).Scan(&listExists)
	if err != nil {
		return false
	}

	return true
}
