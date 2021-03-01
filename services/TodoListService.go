package services

import "database/sql"

// TodoList : TodoList model
type TodoList struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}

// TodoListService : Implementation of TodoListService
type TodoListService struct{}

// GetLists : Returns all lists
func (tls *TodoListService) GetLists(db *sql.DB) ([]TodoList, error) {
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
