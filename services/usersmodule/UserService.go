package usersmodule

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/todo-app-golang/services/usersmodule/dtos"
)

// UserService : Implementation of UserService
type UserService struct{}

// CreateUser : Creates a new user -> It is actually like signing up
func (s *UserService) CreateUser(db *sql.DB, dto *string) (*string, error) {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "insert into users (id, username, password) values (?, ?, ?)"
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
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

// SignIn : Checking username and password -> Returning token
func SignIn(userInfo *dtos.User) (*http.Cookie, error) {

	// Declare the expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims
	claims := &dtos.Claims{
		UserID:   userInfo.ID,
		UserName: userInfo.UserName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(dtos.JwtKey)
	if err != nil {
		return nil, err
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	return cookie, nil
}

// GetUsers : Returns all users
func (s *UserService) GetUsers(db *sql.DB) ([]dtos.User, error) {
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []dtos.User{}

	for rows.Next() {
		var user dtos.User
		if rows.Scan(&user.ID, &user.UserName, &user.Password) != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser : Returns user with given ID
func (s *UserService) GetUser(db *sql.DB, userID string) (*dtos.User, error) {
	user := &dtos.User{}

	q := "select * from users where id = ?"
	err := db.QueryRow(q, userID).
		Scan(&user.ID, &user.UserName, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser : Updates an existing user
func (s *UserService) UpdateUser(db *sql.DB, dto *string) error {

	user := &dtos.User{}
	json.Unmarshal([]byte(*dto), &user)

	// Check whether to be created list is assigning to an existing user
	if !s.doesUserExist(db, &user.ID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "update users set username = ?, password = ? where id = ?"
	stmt, err := db.PrepareContext(ctx, q)
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
func (s *UserService) DeleteUser(db *sql.DB, dto *string) error {

	var userID string
	json.Unmarshal([]byte(*dto), &userID)

	// Check whether to be created list is assigning to an existing user
	if !s.doesUserExist(db, &userID) {
		return nil
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	q := "delete from users where id = ?"
	stmt, err := db.PrepareContext(ctx, q)
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

func (s *UserService) doesUserExist(db *sql.DB, userID *string) bool {
	q := "select id from users where id = ?"
	var userExists string
	err := db.QueryRow(q, userID).Scan(&userExists)
	if err != nil {
		return false
	}

	return true
}
