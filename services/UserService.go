package services

import (
	"database/sql"
)

// User : User model
type User struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"passWord"`
}

// UserService : Implementation of UserService
type UserService struct{}

// GetUsers : Returns all users
func (s *UserService) GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		if rows.Scan(&user.ID, &user.UserName, &user.Password) != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser : Returns user with given ID
func (s *UserService) GetUser(db *sql.DB, userID string) (*User, error) {
	user := &User{}

	q := "select * from users where id = ?"
	err := db.QueryRow(q, userID).
		Scan(&user.ID, &user.UserName, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
