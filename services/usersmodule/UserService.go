package usersmodule

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/todo-app-golang/commons"
	"github.com/todo-app-golang/services/usersmodule/dtos"
)

// UserService : Implementation of UserService
type UserService struct {
	Db *sql.DB
	Cu *commons.CommonUtils
}

// CreateUser : Creates a new user -> It is actually like signing up
func (s *UserService) CreateUser(dto *string) (*string, error) {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "insert into users (id, username, password) values (?, ?, ?)"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &dtos.User{}
	json.Unmarshal([]byte(*dto), &user)

	userID := uuid.New().String()
	res, err := stmt.ExecContext(ctx, userID, user.UserName, user.Password)
	if err != nil {
		return nil, err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return nil, err
	}

	return &userID, nil
}

// GetUser : Returns user with given ID
func (s *UserService) GetUser(userID string) (*dtos.User, error) {
	user := &dtos.User{}

	q := "select * from users where id = ?"
	err := s.Db.QueryRow(q, userID).
		Scan(&user.ID, &user.UserName, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser : Updates an existing user
func (s *UserService) UpdateUser(dto *string) error {

	user := &dtos.User{}
	json.Unmarshal([]byte(*dto), &user)

	// Check whether to be updated user exists
	if !s.Cu.DoesUserExist(s.Db, &user.ID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update users set username = ?, password = ? where id = ?"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.UserName, user.Password, user.ID)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}

// DeleteUser : Deletes an existing user
func (s *UserService) DeleteUser(userID *string) error {

	// Check whether to be deleted user exists
	if !s.Cu.DoesUserExist(s.Db, userID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "delete from users where id = ?"
	stmt, err := s.Db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, userID)
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()
	if numRows != 1 || err != nil {
		return err
	}

	return nil
}

// CheckUser : Checks whether given username and password match
func (s *UserService) CheckUser(userName *string, password *string) (*string, error) {
	q := "select id from users" +
		" where username = ?" +
		" and password = ?"

	var userID string
	err := s.Db.QueryRow(q, userName, password).Scan(&userID)
	if err != nil {
		return nil, errors.New("user_not_found")
	}

	return &userID, nil
}
